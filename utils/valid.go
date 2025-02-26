package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Trans ut.Translator

func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := make(map[string]string)
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, Trans)
		default:
			_ = zh_translations.RegisterDefaultTranslations(v, Trans)
		}
		_ = v.RegisterValidation("nonempty", nonemptyValidation)
		_ = v.RegisterTranslation("nonempty", Trans, func(ut ut.Translator) error {
			return ut.Add("nonempty", "{0}不能为空数组", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("nonempty", fe.Field())
			return t
		})
		return
	}
	return
}

func ValidErr(err error) map[string]string {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"err": err.Error()}
	}
	return RemoveTopStruct(errs.Translate(Trans))
}

func ErrorInfo(err error) map[string]string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {

	}
	return RemoveTopStruct(errs.Translate(Trans))
}
func nonemptyValidation(fl validator.FieldLevel) bool {
	fieldValue := fl.Field()
	switch fieldValue.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		return fieldValue.Len() > 0
	default:
		return false
	}
}
