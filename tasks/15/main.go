package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
  //v и justString указывают на один и тот же участок памяти
  //значит в памяти осталась длинная строка
  fmt.Println(unsafe.StringData(v) == unsafe.StringData(justString)) //true
}

func main() {
  someFunc()
}

*/

var justString string

func someFunc() {
	str := createHugeString(1 << 10)
	justString = strings.Clone(str[:100])
	fmt.Println(unsafe.StringData(str) == unsafe.StringData(justString))
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = '*'
	}
	return string(bytes)
}
