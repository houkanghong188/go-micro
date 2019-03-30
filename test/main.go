package main

import (
	"fmt"
	"go-micro/tool"
)

func main()  {
	fmt.Println(tool.Config["hosts"])
}
