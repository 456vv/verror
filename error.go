package verror
	
import (
	"fmt"
	"reflect"
	"runtime"
//	"unsafe"
	"errors"
)

//设置为 True，可以打印更多错误信息。包括调用函数的名和行号。
var DefaultDebugError bool = true

//错误跟踪
type ErrTrack struct{
	Err 	error
	ErrInfo error
}
func (T *ErrTrack) Error() 	string {if DefaultDebugError {return T.String()};return T.Err.Error()}
func (T *ErrTrack) String() string {return fmt.Sprintf("%v\n%v", T.Err, T.ErrInfo)}


//错误处理（带格式）
//	format string		格式
//	a ...interface{}	错误字段
//	error				错误
func TrackErrorf(format string, a ...interface{}) error {
	return e(2, 0, fmt.Errorf(format, a...))
}

//错误处理
//	a interface{}	传入错误
//	error			返回错误（更多信息）
func TrackError(a interface{}) error {
	switch s := a.(type) {
	case error:
		return e(2, 0, s)
	case string:
		return e(2, 0, fmt.Errorf("verror: %v", s))
	case nil:
		return nil
	default:
		panic(fmt.Sprintf("verror: Error 函数调用，传入错误的数据类型。%s", reflect.TypeOf(a).Kind().String()))
	}
	return nil
}

//错误处理(层次)
//	level, limit int	错误获捕开始层次，限制最多层次（默认无限）
//	err error			传入错误
//	error				返回错误（更多信息）
func ErrorLevel(level, limit int, err error) error {
	return e(level, limit, err)
}

func e(level, limit int, err error) error {
	if err == nil {
		return err
	}
	var et ErrTrack
	et.Err = err
	et.ErrInfo = errors.New("")
	
	for {
		pc, file, line, ok := runtime.Caller(level)
		if !ok {
			break
		}
		runtimeFunc := runtime.FuncForPC(pc)
		et.ErrInfo = fmt.Errorf("%v:%s:%d -> %s\n", et.ErrInfo,  file, line, runtimeFunc.Name())//unsafe.Pointer(runtimeFunc.Entry()),
		level++
		limit--
		if limit == 0 {
			break
		}
	}
	return &et
}
