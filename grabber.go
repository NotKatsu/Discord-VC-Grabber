package main

import (
	"fmt"
	"strings"

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

		authenticationToken += new_encoded_discord_id + "." + strings.Repeat("*", 6) + "." + strings.Repeat("*", 39)

		fmt.Println("Authentication Token: " + authenticationToken)
		fmt.Println("User ID: " + modules.Decode(encoded_discord_id))
	}
}
