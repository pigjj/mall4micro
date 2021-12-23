package log

import (
	"fmt"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/utils"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

func init() {
	createLogsDir()
}

//
// createLogsDir
// @Description: 日志文件夹如果不存在就创建, 如果存在就返回
// @Document:
// @return string
//
func createLogsDir() string {
	projBasePath := utils.ProjectBasePath()
	logDirPath := fmt.Sprintf("%s%slogs", projBasePath, utils.PathSplitFlag)
	if !utils.IsDir(logDirPath) {
		err := os.Mkdir(logDirPath, 0777)
		if err != nil {
			panic(err)
		}
	}
	return logDirPath
}

func InitZapLogger(moduleName string, debug bool) *ZapLogger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(utils.DateTimeFormatStr))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	errorWriter := getWriter(fmt.Sprintf("%s%s%s_error.log", createLogsDir(), utils.PathSplitFlag, moduleName))
	infoWriter := getWriter(fmt.Sprintf("%s%s%s_info.log", createLogsDir(), utils.PathSplitFlag, moduleName))
	debugWriter := getWriter(fmt.Sprintf("%s%s%s_debug.log", createLogsDir(), utils.PathSplitFlag, moduleName))

	var cores = []zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
	}
	if debug {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel))
	}
	// 最后创建具体的Logger
	core := zapcore.NewTee(cores...)

	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	return &ZapLogger{log.Sugar()}
}

//
// getWriter
// @Description: 获取写日志句柄
// @Document:
// @param filename
// @return io.Writer
//
func getWriter(filename string) io.Writer {
	// 生成rotateLogs的Logger 实际生成的文件名 demo.log.YmdH
	// demo.log是指向最新日志的链接
	// 保存90天内的日志，每24小时(整点)分割一次日志
	hook, err := rotateLogs.New(
		fmt.Sprintf("%s-%%Y%%m%%d%%H%%M", filename),
		rotateLogs.WithLinkName(fmt.Sprintf(filename)), // 生成软链，指向最新日志文件
		rotateLogs.WithMaxAge(90*24*time.Hour),         // 文件最大保存时间
		// rotateLogs.WithRotationCount(5),
		// rotateLogs.WithRotationTime(7*time.Hour), // 日志切割时间间隔
		rotateLogs.WithRotationTime(24*time.Hour),  // 日志切割时间间隔
		rotateLogs.WithRotationSize(100*1024*1024), // 在日志切割时间内, 如果日志大小大于此值, 会切换日志文件, 文件流水号+1
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.SugaredLogger.Debug(args...)
}

func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.SugaredLogger.Debugf(template, args...)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.SugaredLogger.Info(args...)
}

func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.SugaredLogger.Infof(template, args...)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.SugaredLogger.Warn(args...)
}

func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.SugaredLogger.Warnf(template, args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.SugaredLogger.Error(args...)
}

func (z *ZapLogger) Errorf(template string, args ...interface{}) {
	z.SugaredLogger.Errorf(template, args...)
}

func (z *ZapLogger) DPanic(args ...interface{}) {
	z.SugaredLogger.DPanic(args...)
}

func (z *ZapLogger) DPanicf(template string, args ...interface{}) {
	z.SugaredLogger.DPanicf(template, args...)
}

func (z *ZapLogger) Panic(args ...interface{}) {
	z.SugaredLogger.Panic(args...)
}

func (z *ZapLogger) Panicf(template string, args ...interface{}) {
	z.SugaredLogger.Panicf(template, args...)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.SugaredLogger.Fatal(args...)
}

func (z *ZapLogger) Fatalf(template string, args ...interface{}) {
	z.SugaredLogger.Fatalf(template, args...)
}
