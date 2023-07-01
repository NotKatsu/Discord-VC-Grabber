package main

import (
	"fmt"

	"github.com/NotKatsu/Discord-VC-Grabber/modules"
)

func main() {
	var discordUserID string

	fmt.Printf("Discord User ID: ")
	fmt.Scanln(&discordUserID)

	result := modules.Encode(string(discordUserID))

	fmt.Println(result)
}
