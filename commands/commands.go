/* The plans for this part of the bot are on hold until the MySQL drivers begin to be used, so custom commands can be
 * made a bit more easily as well as having proper permission checks per command
 */

package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//ExecuteCommand Parses and executes the command from the calling code
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message) {

	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]

	switch msg {
	case "info":
		HandleInfoCommand(s, m)
	case "ping":
		HandlePingCommand(s, m)
	default:
		HandleUnknownCommand(s, m, msg)
	}
}

/*func ExecuteCommandWithArgs() {

}*/

//HandleInfoCommand is the !info command
func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message) {

	message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s```"
	message = fmt.Sprintf(message, "Debug Information", strings.Repeat("-", len("Debug Information")), "ChannelID", m.ChannelID, "Author", m.Author.Username)
	s.ChannelMessageSend(m.ChannelID, message)
}

//HandlePingCommand is for !ping
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {

	s.ChannelMessageSend(m.ChannelID, "pong")
}

//HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
}