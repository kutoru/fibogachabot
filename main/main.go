package main

import (
	"fmt"

	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/tgmanager"
)

func main() {
	fmt.Println("Start")

	glb.LoadEnv()
	glb.InitializeMenuList()

	tgmanager.InitializeBot()
	tgmanager.StartPolling()

	fmt.Println("End")
}
