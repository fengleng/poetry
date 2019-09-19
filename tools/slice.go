/*
@Time : 2019/9/19 19:44
@Author : zxr
@File : slice
@Software: GoLand
*/
package tools

import "reflect"

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
