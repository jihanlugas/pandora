package authentication

import (
	"errors"
	"github.com/jihanlugas/pandora/app/jwt"
	"github.com/jihanlugas/pandora/app/user"
	"github.com/jihanlugas/pandora/config"
	"github.com/jihanlugas/pandora/cryption"
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"github.com/jihanlugas/pandora/utils"
	"time"
)

type AuthenticationUsecase interface {
	SignIn(req *request.Signin) (string, jwt.UserLogin, error)
	RefreshToken(loginUser jwt.UserLogin) (string, error)
	Init(loginUser jwt.UserLogin) (*model.Init, error)
}

type usecaseAuthentication struct {
	repo     Repository
	userRepo user.Repository
}

func (u usecaseAuthentication) SignIn(req *request.Signin) (string, jwt.UserLogin, error) {
	var err error
	var data model.User
	var userLogin jwt.UserLogin

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if utils.IsValidEmail(req.Username) {
		data, err = u.userRepo.GetByEmail(conn, req.Username)
	} else {
		data, err = u.userRepo.GetByUsername(conn, req.Username)
	}

	if err != nil {
		return "", userLogin, err
	}

	err = cryption.CheckAES64(req.Passwd, data.Passwd)
	if err != nil {
		return "", userLogin, errors.New("invalid username or password")
	}

	if !data.IsActive {
		return "", userLogin, errors.New("user not active")
	}

	now := time.Now()
	tx := conn.Begin()

	data.LastLoginDt = &now
	data.UpdateBy = data.ID
	err = u.userRepo.Update(tx, data)
	if err != nil {
		return "", userLogin, err
	}

	err = tx.Commit().Error
	if err != nil {
		return "", userLogin, err
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(config.AuthTokenExpiredHour))

	userLogin.UserID = data.ID
	userLogin.Role = data.Role
	userLogin.PassVersion = data.PassVersion
	token, err := jwt.CreateToken(userLogin, expiredAt)
	if err != nil {
		return "", userLogin, err
	}

	return token, userLogin, err
}

func (u usecaseAuthentication) RefreshToken(loginUser jwt.UserLogin) (string, error) {
	var err error

	expiredAt := time.Now().Add(time.Hour * time.Duration(config.AuthTokenExpiredHour))

	token, err := jwt.CreateToken(loginUser, expiredAt)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u usecaseAuthentication) Init(loginUser jwt.UserLogin) (*model.Init, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	userView, err := u.userRepo.GetViewById(conn, loginUser.UserID)
	if err != nil {
		return nil, err
	}

	init := model.Init{
		UserID:   userView.ID,
		Role:     userView.Role,
		Email:    userView.Email,
		Username: userView.Username,
		NoHp:     userView.NoHp,
		Fullname: userView.Fullname,
		PhotoUrl: "",
	}

	return &init, err
}

func NewAuthenticationUsecase(repo Repository, userRepo user.Repository) AuthenticationUsecase {
	return usecaseAuthentication{
		repo:     repo,
		userRepo: userRepo,
	}
}
