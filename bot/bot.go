package bot

import (
	"fmt" //to print errors

	"drixevel.dev/.devbot/config" //importing our config package which we have created above

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot has been created.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == BotId {
		return
	}

	//Makes sure the message has the prefix in front of it to register it as a command.
	if m.Content[0:1] != config.BotPrefix {
		return
	}

	//Strip the prefix from the message.
	m.Content = m.Content[1:]

	if m.Content == "help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Empty")
	}
}
