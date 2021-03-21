package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gotgbot.com/v0.01/config"
)

var (
	// Logger 日志对象
	Logger    *zap.Logger
	zapConfig zap.Config
)

func init() {
	//读取配置文件中的level设置
	logLevel := config.GetString("log.level")
	//判断是否在debug模式下
	if strings.ToLower(logLevel) == "debug" {
		//debug模式下则原子层级在debuglevel层级下开始
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		//encoder在开发者模式下
		zapConfig.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		//不在debug模式下则原子层级在Infolevel层级下开始
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		//encoder在生产模式下s
		zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	}

	//日志时间戳人类可读
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	//读取配置文件中的log文件设置
	logFile := config.GetString("log.file")
	//如果指明log文件
	if logFile != "" {
		//设置采样策略
		zapConfig.Sampling = &zap.SamplingConfig{
			Initial:    100, //初始值
			Thereafter: 100, //其后值
		}
		zapConfig.Encoding = "json"                    //输出json格式
		zapConfig.OutputPaths = []string{logFile}      //输出文件路径
		zapConfig.ErrorOutputPaths = []string{logFile} //error级别输出文件路径

	} else {
		//如未指明log文件
		zapConfig.OutputPaths = []string{"stderr"}                             //直接标准输出log
		zapConfig.ErrorOutputPaths = []string{"stderr"}                        //直接标准输出error级别log
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //指定颜色输出
		zapConfig.Encoding = "console"                                         //输出格式为console
	}

	Logger, _ = zapConfig.Build() //构建日志
	zap.ReplaceGlobals(Logger)    //将zap计量器核心替换标准log核，注册为全局log
}
