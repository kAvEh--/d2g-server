package main

import (
	"d2g-server/presenter"
	"fmt"
)

func main() {
	rMain := presenter.SetupRouter()
	err := rMain.Run(":8080")
	fmt.Println(err)
}
