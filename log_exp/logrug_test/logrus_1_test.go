// @Author: Perry
// @Date  : 2020/1/13
// @Desc  : 

package logrug_test

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func Test1(t *testing.T) {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
