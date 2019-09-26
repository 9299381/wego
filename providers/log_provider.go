package providers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type LogProvider struct {
}

func (it *LogProvider) Boot() {
	wego.Config("log", &configs.LogConfig{})
}

func (it *LogProvider) Register() {
	logger := logrus.New()
	if args.Mode == "prod" {
		//写入文件
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		logger.SetOutput(src)
		//设置日志级别
		logger.SetLevel(logrus.InfoLevel)
		logger.AddHook(it.getLogHook())

	} else {
		src := os.Stdout
		logger.SetOutput(src)
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: constants.YmdHis,
		})
		logger.SetLevel(logrus.DebugLevel)

	}
	wego.App.Logger = logger

}
func (it *LogProvider) getLogHook() *lfshook.LfsHook {
	config := wego.Config("log").(*configs.LogConfig)
	logFilePath := config.LogFilePath
	logFileName := config.LogFileName
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	logWriter, _ := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	return lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: constants.YmdHis,
	})
}
