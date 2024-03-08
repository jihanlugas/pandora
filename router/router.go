package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jihanlugas/pandora/app/app"
	"github.com/jihanlugas/pandora/app/authentication"
	"github.com/jihanlugas/pandora/app/district"
	"github.com/jihanlugas/pandora/app/jwt"
	"github.com/jihanlugas/pandora/app/ktp"
	"github.com/jihanlugas/pandora/app/log"
	"github.com/jihanlugas/pandora/app/province"
	"github.com/jihanlugas/pandora/app/regency"
	"github.com/jihanlugas/pandora/app/user"
	"github.com/jihanlugas/pandora/app/village"
	"github.com/jihanlugas/pandora/config"
	"github.com/jihanlugas/pandora/constant"
	"github.com/jihanlugas/pandora/db"
	_ "github.com/jihanlugas/pandora/docs"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/response"
	"github.com/jihanlugas/pandora/scheduler"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	echoSwagger "github.com/swaggo/echo-swagger"
	"io"
	"net/http"
)

func Init() *echo.Echo {
	var err error
	router := websiteRouter()

	provinceRepo := province.NewRepository()
	regencyRepo := regency.NewRepository()
	districtRepo := district.NewRepository()
	villageRepo := village.NewRepository()
	authenticationRepo := authentication.NewRepository()
	userRepo := user.NewRepository()
	ktpRepo := ktp.NewRepository()

	provinceUsecase := province.NewUsecase(provinceRepo)
	regencyUsecase := regency.NewUsecase(regencyRepo)
	districtUsecase := district.NewUsecase(districtRepo)
	villageUsecase := village.NewUsecase(villageRepo)
	authenticationUsecase := authentication.NewAuthenticationUsecase(authenticationRepo, userRepo)
	userUsecase := user.NewUsecase(userRepo)
	ktpUsecase := ktp.NewUsecase(ktpRepo)

	provinceHandler := province.Handler(provinceUsecase)
	regencyHandler := regency.Handler(regencyUsecase)
	districtHandler := district.Handler(districtUsecase)
	villageHandler := village.Handler(villageUsecase)
	authenticationHandler := authentication.NewAuthenticationHandler(authenticationUsecase)
	userHandler := user.Handler(userUsecase)
	ktpHandler := ktp.Handler(ktpUsecase)

	router.Use(loggerMiddleware)

	router.GET("/swg/*", echoSwagger.WrapHandler)
	router.GET("/", app.Ping)

	router.POST("/sign-in", authenticationHandler.SignIn)
	router.GET("/sign-out", authenticationHandler.SignOut)
	//router.POST("/sign-up", authenticationHandler.SignUp)
	router.GET("/refresh-token", authenticationHandler.RefreshToken, checkTokenMiddleware)
	router.GET("/init", authenticationHandler.Init, checkTokenMiddleware)

	provinceRouter := router.Group("/province")
	provinceRouter.GET("/:id", provinceHandler.GetById)
	provinceRouter.GET("/page", provinceHandler.Page)
	provinceRouter.GET("/list", provinceHandler.List)

	regencyRouter := router.Group("/regency")
	regencyRouter.GET("/:id", regencyHandler.GetById)
	regencyRouter.GET("/page", regencyHandler.Page)
	regencyRouter.GET("/list", regencyHandler.List)

	districtRouter := router.Group("/district")
	districtRouter.GET("/:id", districtHandler.GetById)
	districtRouter.GET("/page", districtHandler.Page)
	districtRouter.GET("/list", districtHandler.List)

	villageRouter := router.Group("/village")
	villageRouter.GET("/:id", villageHandler.GetById)
	villageRouter.GET("/page", villageHandler.Page)
	villageRouter.GET("/list", villageHandler.List)

	userRouter := router.Group("/user")
	userRouter.GET("/:id", userHandler.GetById)
	userRouter.POST("", userHandler.Create, checkTokenAdminMiddleware)
	userRouter.PUT("/:id", userHandler.Update, checkTokenAdminMiddleware)
	userRouter.DELETE("/:id", userHandler.Delete, checkTokenAdminMiddleware)
	userRouter.GET("/page", userHandler.Page, checkTokenAdminMiddleware)
	userRouter.POST("/change-password", userHandler.ChangePassword, checkTokenMiddleware)

	ktpRouter := router.Group("/ktp")
	ktpRouter.GET("/:id", ktpHandler.GetById)
	ktpRouter.POST("", ktpHandler.Create, checkTokenMiddleware)
	ktpRouter.PUT("/:id", ktpHandler.Update, checkTokenMiddleware)
	ktpRouter.DELETE("/:id", ktpHandler.Delete, checkTokenMiddleware)
	ktpRouter.GET("/page", ktpHandler.Page, checkTokenMiddleware)

	schedule := cron.New()
	defer schedule.Stop()

	specDeleteLog := config.ScheduleDeleteLog
	_, err = schedule.AddFunc(specDeleteLog, scheduler.DeleteLog)
	if err != nil {
		fmt.Printf("Error when start schedule delete log %s\n", err.Error())
		panic(err)
	}

	go schedule.Start()

	return router

}

func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		body, _ := io.ReadAll(c.Request().Body)
		c.Set(constant.Request, string(body))
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// Call next handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		res := ""
		if c.Get(constant.Response) != nil {
			res = string(c.Get(constant.Response).([]byte))
		}

		loginUserString := ""
		loginUser, err := jwt.GetUserLoginInfo(c)
		if err == nil {
			loginUserByte, _ := json.Marshal(loginUser)
			loginUserString = string(loginUserByte)
		}

		logRepo := log.NewRepository()
		logUsecase := log.NewUsecase(logRepo)

		logData := model.Log{
			ClientIP:  c.Request().RemoteAddr,
			Method:    c.Request().Method,
			Path:      c.Request().URL.String(),
			Code:      c.Response().Status,
			Loginuser: loginUserString,
			Request:   string(body),
			Response:  res,
		}

		_ = logUsecase.Create(logData)

		return nil
	}
}

func httpErrorHandler(err error, c echo.Context) {
	var errorResponse *response.Response
	code := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		// Handle pada saat URL yang di request tidak ada. atau ada kesalahan server.
		code = e.Code
		errorResponse = &response.Response{
			Status:  false,
			Message: fmt.Sprintf("%v", e.Message),
			Payload: map[string]interface{}{},
			Code:    code,
		}
	case *response.Response:
		errorResponse = e
	default:
		// Handle error dari panic
		code = http.StatusInternalServerError
		if config.Debug {
			errorResponse = &response.Response{
				Status:  false,
				Message: err.Error(),
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		} else {
			errorResponse = &response.Response{
				Status:  false,
				Message: "Internal server error",
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		}
	}

	js, err := json.Marshal(errorResponse)
	if err == nil {
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, js)
	} else {
		b := []byte("{error: true, message: \"unresolved error\"}")
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
	}
}

func checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		userLogin, err := jwt.ExtractClaims(c.Request().Header.Get(config.HeaderAuthName))
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, err.Error(), response.Payload{}).SendJSON(c)
		}

		conn, closeConn := db.GetConnection()
		defer closeConn()

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
		}

		if user.PassVersion != userLogin.PassVersion {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired~", response.Payload{}).SendJSON(c)
		}

		c.Set(constant.TokenUserContext, userLogin)
		return next(c)
	}
}

func checkTokenAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		userLogin, err := jwt.ExtractClaims(c.Request().Header.Get(config.HeaderAuthName))
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, err.Error(), response.Payload{}).SendJSON(c)
		}

		conn, closeConn := db.GetConnection()
		defer closeConn()

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
		}

		if user.Role != constant.RoleAdmin {
			return response.ErrorForce(http.StatusUnauthorized, "permission denied.", response.Payload{}).SendJSON(c)
		}

		if user.PassVersion != userLogin.PassVersion {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired~", response.Payload{}).SendJSON(c)
		}

		c.Set(constant.TokenUserContext, userLogin)
		return next(c)
	}
}
