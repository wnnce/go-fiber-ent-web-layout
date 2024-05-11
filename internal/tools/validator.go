package tools

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var defaultStructValidator *StructValidator

type StructValidator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (s *StructValidator) Engine() any {
	return s.validate
}

// ValidateStruct 验证结构体，将保修信息转化为可读的错误信息返回
func (s *StructValidator) ValidateStruct(out any) error {
	err := s.validate.Struct(out)
	if err == nil {
		return nil
	}
	var validateErrors validator.ValidationErrors
	if !errors.As(err, &validateErrors) || len(validateErrors) == 0 {
		return FiberRequestError("Struct parameter error")
	}
	return FiberRequestError(validateErrors[0].Translate(s.trans))
}

func init() {
	defaultStructValidator = newStruckValidator()
}

// 初始化结构体验证
func newStruckValidator() *StructValidator {
	zhTranslator := zh.New()
	uni := ut.New(zhTranslator, zhTranslator)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	return &StructValidator{
		validate: validate,
		trans:    trans,
	}
}

func DefaultStructValidator() *StructValidator {
	return defaultStructValidator
}
