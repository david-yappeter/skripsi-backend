package main

import (
	"fmt"
	"myapp/cmd"
)

// @title		Mortal Health - Clinic Pilot API
// @version	0.0.1
// @host		cp-api.mortalhealth.com
// @BasePath	/
func main() {
	temp := true
	if temp {
		fmt.Println("TEMP")

		temp = false
	} else {
		fmt.Println("OTHER")
	}

	return

	cmd.Execute()
}
