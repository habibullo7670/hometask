package main

import (
	"fmt"
	"strings"
	"strconv"

)

func PrintMainScreen() {
	fmt.Println("Welcome to our pay terminal!")
	fmt.Println("Choose your operator:")
	fmt.Println("1 - Megafon")
	fmt.Println("2 - Babilon")
	fmt.Println("3 - Tcell")
	fmt.Println("4 - Beeline")
}

func CheckNumber(operatorID int, number string) bool {
	var checkPrefix bool
	operatorPreffix := map[int]string{
		1: "90,882",
		2: "918,98",
		3: "93,55",
		4: "919,917,916,915,914,913,912,911,910",
	}

	if len(number) != 9 {
		return false
	}
	_,err:=strconv.Atoi(number)
	if err!=nil {
		return false
	}

	prefix := operatorPreffix[operatorID]
	arrPrefix := strings.Split(prefix, ",")

	for _, _prefixValue := range arrPrefix {
		checkPrefix = strings.HasPrefix(number, _prefixValue)

		if checkPrefix {
			return true
		}
	    
	}
	return checkPrefix
}

