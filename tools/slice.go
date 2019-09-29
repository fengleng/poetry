/*
@Time : 2019/9/19 19:44
@Author : zxr
@File : slice
@Software: GoLand
*/
package tools

import (
	"math/rand"
	"reflect"
	"time"
)

//检查变量是否在slice中
func IsExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

//生成随机整数slice
func RandInt64Slice(len, maxNumber int) []int64 {
	var intSlice = make([]int64, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intSlice[i] = rand.Int63n(int64(maxNumber))
	}
	return intSlice
}
