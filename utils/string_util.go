package utils

import "strings"

//去空格 去换行
func FilterString(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}

//去空格
func FilterSpace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	return s
}

//去换行
func FilterEnter(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	return s
}