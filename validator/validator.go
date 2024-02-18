package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jihanlugas/pandora/app/user"
	"github.com/jihanlugas/pandora/config"
	"github.com/jihanlugas/pandora/db"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var (
	Validate *CustomValidator
	regxNoHp *regexp.Regexp
	regExt   *regexp.Regexp
)

type CustomValidator struct {
	validator *validator.Validate
}

func init() {
	Validate = NewValidator()
	regxNoHp = regexp.MustCompile(`((^\+?628\d{8,14}$)|(^0?8\d{8,14}$)){1}`)
	regExt = regexp.MustCompile(`(?i)^\.?(jpe?g|png|webp|)$`)
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// ValidateVar for validate field against tag. Expl: ValidateVar("abc@gmail.com", "required,email")
func (v *CustomValidator) ValidateVar(field interface{}, tag string) error {
	return v.validator.Var(field, tag)
}

func NewValidator() *CustomValidator {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = validate.RegisterValidation("notexists", notExistsOnDbTable)
	_ = validate.RegisterValidation("existsdata", existsDataOnDbTable)
	_ = validate.RegisterValidation("no_hp", validNoHp)
	_ = validate.RegisterValidation("passwdComplex", checkPasswordComplexity)
	_ = validate.RegisterValidation("photo", photoCheck, true)
	//_ = validate.RegisterValidation("hiragana", hiragana)
	//_ = validate.RegisterValidation("katakana", katakana)
	//_ = validate.RegisterValidation("kana", kana)
	//_ = validate.RegisterValidation("kanji", kanji)
	//_ = validate.RegisterValidation("electionTypeProvince", electionTypeProvince)
	//_ = validate.RegisterValidation("electionTypeRegency", electionTypeRegency)
	//_ = validate.RegisterValidation("electionTypeDistrictdapil", electionTypeDistrictdapil)

	return &CustomValidator{
		validator: validate,
	}
}

func notExistsOnDbTable(fl validator.FieldLevel) bool {
	var err error
	params := strings.Fields(fl.Param())

	userRepo := user.NewRepository()

	conn, closeConn := db.GetConnection()
	defer closeConn()

	switch params[0] {
	case "username":
		username := strings.TrimSpace(fl.Field().String())
		if username == "" {
			return true
		}
		_, err = userRepo.GetByUsername(conn, username)
		if err != nil && err == gorm.ErrRecordNotFound {
			return true
		}
		return false

	case "email":
		email := strings.TrimSpace(fl.Field().String())
		if email == "" {
			return true
		}

		_, err = userRepo.GetByEmail(conn, email)
		if err != nil && err == gorm.ErrRecordNotFound {
			return true
		}
		return false

	case "no_hp":
		noHp := strings.TrimSpace(fl.Field().String())
		if noHp == "" {
			return true
		}

		_, err = userRepo.GetByNoHp(conn, noHp)
		if err != nil && err == gorm.ErrRecordNotFound {
			return true
		}
		return false

	}

	return false
}

func existsDataOnDbTable(fl validator.FieldLevel) bool {
	var err error
	params := strings.Fields(fl.Param())

	//if fl.Field().Int() == 0 {
	//	return true
	//}

	userRepo := user.NewRepository()

	conn, closeConn := db.GetConnection()
	defer closeConn()

	switch params[0] {
	case "user_id":
		ID := fl.Field().String()
		if ID == "" {
			return true
		}
		_, err = userRepo.GetById(conn, ID)
		if err != nil {
			return false
		}
		return true
		//case "company_id":
		//	ID := fl.Field().String()
		//	if ID == "" {
		//		return true
		//	}
		//	_, err = companyRepo.GetById(conn, ID)
		//	if err != nil {
		//		return false
		//	}
		//	return true
	}
	return false
}

func IsSameDate(date1, date2 *time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func checkPasswordComplexity(fl validator.FieldLevel) bool {
	passwd := fl.Field().String()

	var capitalFlag, lowerCaseFlag, numberFlag bool
	for _, c := range passwd {
		if unicode.IsUpper(c) {
			capitalFlag = true
		} else if unicode.IsLower(c) {
			lowerCaseFlag = true
		} else if unicode.IsDigit(c) {
			numberFlag = true
		}
		if capitalFlag && lowerCaseFlag && numberFlag {
			return true
		}
	}
	return false
}

func validNoHp(fl validator.FieldLevel) bool {
	return regxNoHp.MatchString(fl.Field().String())
}

func photoCheck(fl validator.FieldLevel) bool {
	params := strings.Fields(fl.Param())

	if len(params) == 0 {
		return true
	}
	parentVal := fl.Parent()
	if parentVal.Kind() == reflect.Ptr {
		parentVal = reflect.Indirect(parentVal)
	}
	// field photo harus dengan tipe data: *multipart.FileHeader ( pointer )
	photoVal := parentVal.FieldByName(params[0])
	if photoVal.Kind() != reflect.Ptr {
		return false
	}
	if photoVal.IsZero() {
		return true
	}
	if f, ok := photoVal.Interface().(*multipart.FileHeader); !ok {
		return false
	} else {
		if !regExt.MatchString(filepath.Ext(f.Filename)) {
			return false
		}
		if f.Size > config.MaxSizeUploadPhotoByte {
			return false
		}
		return true
	}
}
