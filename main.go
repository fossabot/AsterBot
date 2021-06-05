package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var config map[string]interface{}

func main() {

	//Read Configuration from file
	config = GetConfig("./conf.json")

	dg, err := discordgo.New("Bot " + fmt.Sprintf("%v", config["token"]))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func GetConfig(filename string) map[string]interface{} {
	file, _ := ioutil.ReadFile(filename)
	var config map[string]interface{}
	json.Unmarshal(file, &config)
	return config
}
