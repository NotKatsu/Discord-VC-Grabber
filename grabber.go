package main

import (
	"fmt"

	"github.com/NotKatsu/Discord-VC-Grabber/modules"
)

func main() {
	for {
		var discordUserID string
		var authenticationToken string

		fmt.Printf("Discord User ID: ")
		fmt.Scanln(&discordUserID)

		encoded_discord_id := modules.Encode(string(discordUserID))
		new_encoded_discord_id := encoded_discord_id[:len(encoded_discord_id)-2]

		authenticationToken += new_encoded_discord_id + "."
		fmt.Println(authenticationToken)
	}
}
