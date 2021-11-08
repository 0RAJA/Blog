package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() (ret []string) {
	for _, v := range v {
		ret = append(ret, v.Error())
	}
	return
}

func (v *ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

/*
BindAndValid
在上述代码中，我们主要是针对入参校验的方法进行了二次封装，
在 BindAndValid 方法中，通过 ShouldBind 进行参数绑定和入参校验，
当发生错误后，再通过上一步在中间件 Translations 设置的 Translator 来对错误消息体进行具体的翻译行为。
*/
func BindAndValid(c *gin.Context, v interface{}) (ok bool, errs ValidErrors) {
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for k, v := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, errs
	}
	return true, nil
}
