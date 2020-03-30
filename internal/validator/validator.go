package validator

import vv9 "gopkg.in/go-playground/validator.v9"

func Validate(t interface{}) error {
	v := vv9.New()
	v.RegisterAlias("nonnil", "required")
	return v.Struct(t)
}
