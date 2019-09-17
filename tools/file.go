/*
@Time : 2019/9/17 14:33
@Author : zxr
@File : file
@Software: GoLand
*/
package tools

import "os"

func FileExists(file string) (ret bool, err error) {
	if _, err := os.Stat(file); err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
