package validator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"

	"mall/internal/pkg/constant"
)

var Trans ut.Translator

func RegisterTranslation(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhTrans := zh.New()
		enTrans := en.New()

		uniTrans := ut.New(enTrans, zhTrans, enTrans)
		Trans, ok = uniTrans.GetTranslator(locale)
		if !ok {
			err = fmt.Errorf("uniTrans.GetTranslator(%s) failed", locale)
			return errors.Wrap(err, constant.TransGetTranslatorError)
		}

		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}
