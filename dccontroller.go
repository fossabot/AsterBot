package main //discord controller

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToUpper(m.Content)
	if Contain(content, "ASTER") && (m.Author.ID != fmt.Sprintf("%v", config["bot_id"])) { // match with bot id because bot also read its own replies.
		checkMod := CheckMod(m.Author.ID, strings.Split(fmt.Sprintf("%v", config["mod"]), ", "))

		// Mods only Commands
		if checkMod {
			// Lock Mode
			if Contain(content, "LOCK") {
				if Contain(content, "ON") {
					lock_mode = true
					Send(s, m, "Lock Mode On 👻, \n")
				}
				if Contain(content, "OFF") {
					lock_mode = false
					Send(s, m, "Lock Mode Off 😗, \n")
				}
			}

			// Container Stop
			if Contain(content, "STOP") {
				ContainerStop(fmt.Sprintf("%v", config["container_id"]))
				Send(s, m, "Stopping Server 😔, \n")
			}

			// Container Restart
			if Contain(content, "RESTART") {
				ContainerRestart(fmt.Sprintf("%v", config["container_id"]))
				Send(s, m, "Restarting Server 😉, \n")
			}

			// Whitelist list
			if Contain(content, "WHITELIST") && Contain(content, " LIST") {
				if ContainerExec(fmt.Sprintf("%v", config["container_id"]), "whitelist list") {
					status, out := ContainerLog(fmt.Sprintf("%v", config["container_id"]), 0)
					if status {
						out = out[4 : len(out)-6]
						Send(s, m, "Result:\n"+GetWhitelist(out))
					}
				}
			}

			// Whitelist Add
			if name, check := GetName(m.Content); Contain(content, "WHITELIST") && Contain(content, " ADD") && check {
				if ContainerExec(fmt.Sprintf("%v", config["container_id"]), "whitelist add "+name) {
					status, out := ContainerLog(fmt.Sprintf("%v", config["container_id"]), 0)
					if status {
						Send(s, m, "Result:\n"+out)
					}
				}
			}

			// Whitelist Remove
			if name, check := GetName(m.Content); Contain(content, "WHITELIST") && Contain(content, " REMOVE") && check {
				if ContainerExec(fmt.Sprintf("%v", config["container_id"]), "whitelist remove "+name) {
					status, out := ContainerLog(fmt.Sprintf("%v", config["container_id"]), 0)
					if status {
						Send(s, m, "Result:\n"+out)
					}
				}
			}

		}

		// Public Commands
		if !lock_mode || checkMod {
			// Container Start
			if Contain(content, "START") {
				if ContainerStatus(fmt.Sprintf("%v", config["container_id"])) {
					Send(s, m, "Server is running already 👀, \n")
				} else {
					if ContainerStart(fmt.Sprintf("%v", config["container_id"])) {
						Send(s, m, "Starting Server 😍, \n")
					}
				}
			}

			// Container Status
			if Contain(content, "STATUS") && !lock_mode {
				if ContainerStatus(fmt.Sprintf("%v", config["container_id"])) {
					Send(s, m, "Server is Up 😎, \n")
				} else {
					Send(s, m, "Server is Down 😓, \n")
				}
			}

			// About Bot
			if Contain(content, "ABOUT") && !lock_mode {
				Send(s, m, "Server Status ကိုကြည့်နိုင်အောင် PeterZam ကရေးထားတာပါ။\nSource Code : https://github.com/peterzam/asterianbot\nAvaliable Commands:\n Start, Stop, Restart, MT, Status, About")
			}

			// List Player
			if Contain(content, "LIST") && !Contain(content, "WHITELIST") && !lock_mode {
				if ContainerExec(fmt.Sprintf("%v", config["container_id"]), "list") {
					status, out := ContainerLog(fmt.Sprintf("%v", config["container_id"]), 0)
					if status {
						Send(s, m, "Result:\n"+out)
					}
				}

			}
		}

	}
}

// Shortened strings.Contains
func Contain(content string, sub string) bool {
	return strings.Contains(content, sub)
}

// Shortened s.ChannelMessageSend
func Send(s *discordgo.Session, m *discordgo.MessageCreate, str string) {
	s.ChannelMessageSend(m.ChannelID, str+m.Author.Mention())
}

// Get Name String between "[" and "]"
func GetName(str string) (result string, found bool) {
	s := strings.Index(str, "[")
	if s == -1 {
		return result, false
	}
	newS := str[s+len("["):]
	e := strings.Index(newS, "]")
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

// Match author id whit mod id (check if its mods or not)
func CheckMod(id string, str []string) bool {
	for _, str_id := range str {
		if str_id == id {
			return true
		}
	}
	return false
}

// Parse docker output json to string name
func GetWhitelist(str string) string {
	value := gjson.Get(str, "result.#.name")
	result := strings.ReplaceAll(value.String(), ",", ",\n")
	result = result[1:len(result)-1] + "\n"
	return result
}
