package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

// 日志等级
const (
	LevelDebug 	Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// 根据常量返回对应字符串
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// 日志结构
type Logger struct {
	newLogger	*log.Logger
	ctx 		context.Context
	level		Level
	fields 		Fields
	callers 	[]string
}

// Logger 实例化
func NewLogger(w io.Writer,prefix string,flag int) *Logger {
	l := log.New(w,prefix,flag)
	return &Logger{newLogger: l}
}

// 日志克隆
func (l *Logger) clone() *Logger {
	n1 := *l
	return &n1
}

// 设置日志等级
func (l *Logger) WithLevel(newLevel Level) *Logger {
	newLogger := l.clone()
	newLogger.level = newLevel
	return newLogger
}

// 设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	newLogger := l.clone()
	if newLogger.fields == nil {
		newLogger.fields = make(Fields)
	}
	for k, v := range f {
		newLogger.fields[k] = v
	}
	return newLogger
}

// 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// 设置当前某一层调用栈的信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc,file,line,ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s",file,line,f.Name())}
	}
	return ll
}

// 设置当前的整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr,maxCallerDepth)
	depth := runtime.Callers(minCallerDepth,pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame,more := frames.Next();more;frame,more = frames.Next() {
		callers = append(callers,fmt.Sprintf("%s: %d %s",frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// 日志格式化
func (l *Logger) JSONFormat(message string) map[string]interface{} {
	data := make(Fields,len(l.fields)+4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

// 日志输出
func (l *Logger) Output(message string) {
	body,_ := json.Marshal(l.JSONFormat(message))
	content := string(body)
	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Print(content)
	case LevelPanic:
		l.newLogger.Print(content)
	}
}

// 日志分级输出
func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format,v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string,v ...interface{})  {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format,v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format,v...))
}