package app

import (
	"gin-blog/utils/logging"
	"github.com/beego/beego/v2/core/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
