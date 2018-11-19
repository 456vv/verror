# vmap [![Build Status](https://travis-ci.org/456vv/verror.svg?branch=master)](https://travis-ci.org/456vv/verror)
golang verror，支持打印出详细的错误信息。

# **列表：**
```go
var DefaultDebugError bool = true																		// 设置为 True，可以打印更多错误信息。包括调用函数的名和行号。
type ErrTrack struct{																			// 错误跟踪
	Err 	error																						// 原始错误
	ErrInfo error																						// 错误行
}
func (T *ErrTrack) Error() string  		                                                                // 错误
func (T *ErrTrack) String() string                                                  					// 字符
func TrackErrorf(format string, a ...interface{}) error													// 错误处理（带格式）
func TrackError(a interface{}) error																	// 错误处理
func ErrorLevel(level, limit int, err error) error														// 错误处理(层次)
```