package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

//GetValidate 初始化验证器
func GetValidate() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}
