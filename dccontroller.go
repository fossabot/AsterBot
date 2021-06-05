package main //discord controller

// fmt.Println(CheckStatus("5c8de1057d2d"))
// fmt.Println(StartContainer("5c8de1057d2d"))
// fmt.Println(StopContainer("5c8de1057d2d"))
// fmt.Println(RestartContainer("5c8de1057d2d"))

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToUpper(m.Content)
	if strings.Contains(content, "ASTER") && (m.Author.ID != fmt.Sprintf("%v", config["bot_id"])) {
		for _, mod_id := range strings.Split(fmt.Sprintf("%v", config["mod"]), ", ") {
			if m.Author.ID == mod_id {
				if strings.Contains(content, "STOP") {
					ContainerStop(fmt.Sprintf("%v", config["container_id"]))
					s.ChannelMessageSend(m.ChannelID, "Stopping Server 😔, \n"+m.Author.Mention())
				}
				if strings.Contains(content, "RESTART") {
					ContainerRestart(fmt.Sprintf("%v", config["container_id"]))
					s.ChannelMessageSend(m.ChannelID, "Restarting Server 😉, \n"+m.Author.Mention())
				}
				if strings.Contains(content, "MT") {
					if strings.Contains(content, "ON") {
						maintainance_mode = true
						s.ChannelMessageSend(m.ChannelID, "Maintainance Mode On 👻, \n"+m.Author.Mention())
					}
					if strings.Contains(content, "OFF") {
						maintainance_mode = false
						s.ChannelMessageSend(m.ChannelID, "Maintainance Mode Off 😗, \n"+m.Author.Mention())
					}
				}
			}
		}

		if strings.Contains(content, "START") && !maintainance_mode {
			if ContainerStatus(fmt.Sprintf("%v", config["container_id"])) {
				s.ChannelMessageSend(m.ChannelID, "Server is running already 👀, \n"+m.Author.Mention())
			} else {
				if ContainerStart(fmt.Sprintf("%v", config["container_id"])) {
					s.ChannelMessageSend(m.ChannelID, "Starting Server 😍, \n"+m.Author.Mention())
				}
			}
		}

		if strings.Contains(content, "STATUS") && !maintainance_mode {
			if ContainerStatus(fmt.Sprintf("%v", config["container_id"])) {
				s.ChannelMessageSend(m.ChannelID, "Server is Up 😎, \n"+m.Author.Mention())
			} else {
				s.ChannelMessageSend(m.ChannelID, "Server is Down 😓, \n"+m.Author.Mention())
				m.Author.Mention()
			}
		}

		if strings.Contains(content, "ABOUT") {
			s.ChannelMessageSend(m.ChannelID, "Server Status ကိုကြည့်နိုင်အောင် PeterZam ကရေးထားတာပါ။\nSource Code : https://github.com/peterzam/asterianbot\nAvaliable Commands:\n Start, Stop, Restart, MT, Status, About")
		}
	}
}
