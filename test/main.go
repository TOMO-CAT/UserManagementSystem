package main

import (
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/TOMO-CAT/UserManagementSystem/proto/config"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(strconv.FormatUint(rand.Uint64(), 10))
	}
}
