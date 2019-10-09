/*
@Time : 2019/9/17 14:33
@Author : zxr
@File : file
@Software: GoLand
*/
package tools

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

func FileExists(file string) (ret bool, err error) {
	if _, err := os.Stat(file); err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

//得到文件指针
func FilePointer(fileName string) (file *os.File, err error) {
	if file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		return
	}
	if file != nil {
		return file, nil
	}
	return nil, errors.New("file point error")
}

//写入recover到日志文件
func WriteRecover(err interface{}) {
	if file, e := FilePointer("logs/error.log"); e == nil {
		errStr := fmt.Errorf("%v", err).Error()
		for i := 0; i <= 5; i++ {
			_, runFile, line, ok := runtime.Caller(i)
			if ok {
				errStr += fmt.Sprintf("  File:%s,  Line:%d\n", runFile, line)
			} else {
				errStr += "\n"
				break
			}
		}
		_, _ = file.WriteString(errStr)
		_ = file.Close()
	}
	return
}
