package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type IValidation interface {
	ValidateStruct(req interface{}) error
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation(val *validator.Validate) IValidation {
	return &Validation{
		validate: val,
	}
}

func (v *Validation) ValidateStruct(req interface{}) error {
	err := v.validate.Struct(req)
	if err != nil {
		return v.InvalidValidationError(err)
	}
	return nil
}

func (v *Validation) InvalidValidationError(err error) error {
	var errMessages []string
	for _, e := range err.(validator.ValidationErrors) {
		// 根据不同的验证标签生成对应的错误信息
		var msg string
		switch e.Tag() {
		case "required":
			msg = fmt.Sprintf("%s为必填字段，请检查", e.Field())
		case "email":
			msg = fmt.Sprintf("%s格式无效，请检查", e.Field())
		case "min":
			msg = fmt.Sprintf("%s长度最少需要%s个字符", e.Field(), e.Param())
		default:
			msg = fmt.Sprintf("%s字段验证失败（%s）", e.Field(), e.Tag())
		}
		errMessages = append(errMessages, msg)
	}

	// 合并所有错误信息
	fullMessage := "validation error：" + strings.Join(errMessages, "\n")
	return errors.New(fullMessage)
}
