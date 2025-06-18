package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj any) string {
	//global.Log.Errorf("错误：%v", err.Error())
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				if msg != "" {
					return msg
				}
			}
		}
	}
	return err.Error()
}
