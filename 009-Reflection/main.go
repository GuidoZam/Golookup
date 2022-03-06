package main

import (
	"fmt"
    "reflect"
    "runtime"
    "strings"
)

func GetFunctionName(temp interface{}) string {
    strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
    return strs[len(strs)-1]
}

func ThisIsAnEmptyFunction() {
}

func main() {
	fmt.Println(GetFunctionName(main))
	fmt.Println(GetFunctionName(ThisIsAnEmptyFunction))
}