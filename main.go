package main

import (
	"fmt"
	"strconv"

	"github.com/UncleMutu/simple-server/entity"
)

func main() {
	user := entity.User{
		ID:   0,
		Name: "testing name user",
	}

	fmt.Println("testing print struct from vim, ", user)

	testingString, err := strconv.Atoi("10")
	if err != nil {
		panic(err)
	}
	fmt.Println("testing string: ", testingString)
}
