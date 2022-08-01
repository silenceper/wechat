package wechat

import (
	"github.com/sirupsen/logrus"
	"io"
)

type LogConfig struct {
	Formatter logrus.Formatter
	Output    io.Writer
	Level     logrus.Level
}
