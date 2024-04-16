package tools

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate           *validator.Validate
	trans              ut.Translator
	customErrorMessage = map[string]string{} // 自定义的错误消息
)

// init 自动初始化参数校验类和翻译器
// 翻译器添加自定义的错误信息
func init() {
	zhTranslator := zh.New()
	uni := ut.New(zhTranslator, zhTranslator)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	for k, v := range customErrorMessage {
		_ = validate.RegisterTranslation(k, trans, func(ut ut.Translator) error {
			return ut.Add(k, fmt.Sprintf("{0} %s", v), true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(k, fe.Field())
			return t
		})
	}
}

func StructFieldValidation(entity interface{}) string {
	err := validate.Struct(entity)
	if err == nil {
		return ""
	}
	var validateErrors validator.ValidationErrors
	if !errors.As(err, &validateErrors) || len(validateErrors) == 0 {
		return "Request parameter error"
	}
	return validateErrors[0].Translate(trans)
}
