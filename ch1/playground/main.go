package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strInt := "123"
	anInt, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d\n", anInt)

	array := [4]string{"b", "r", "a", "d"}

	fmt.Printf("%s\n", array)

	str := strings.Join(array[:], "")

	fmt.Printf("%s\n", str)
}
