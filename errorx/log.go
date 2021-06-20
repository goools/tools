package errorx

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewErrorWithLog(format string, args ...interface{}) error {
	logrus.Errorf(format, args...)
	return fmt.Errorf(format, args...)
}
