package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s\n",&str)
	var result = coding(str)
	for _, c :=range result{
		fmt.Print("",c)
	}
}
func coding(str string) []string {
	var replay = 0
	value:= make([]string, len(str))
	var nextgen = 0
	for i := 0;i<len(str);i++{
		if i != 0 && i<len(str){
			if str[i] == str[i-1]{replay++} else {
				value[nextgen] = string(str[i-1]) + strconv.Itoa(replay)
				nextgen++
				replay=1
				continue
			}
		}else if i!=len(str){
			replay++
			continue
		}
		value[nextgen] = string(str[i]) + strconv.Itoa(replay)
	}
	returning := value[:nextgen+1]
	return returning
}
