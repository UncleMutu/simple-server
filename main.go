package main

import (
	"fmt"

	"github.com/UncleMutu/simple-server/entity"
)

func main() {
	// fmt.Println("testing aja")
	// numb := int64(10)
	// strconvert := strconv.Itoa(int(numb))
	// fmt.Println("string convert from integer: ", strconvert)

	user := entity.User{
		ID:   0,
		Name: "testing name user",
	}

	fmt.Println("testing print struct from vim, ", user)
}
