package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string
	Guild string
	User  string
}

func main() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
		return
	}

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	roleA := "866734447842557992"
	roleB := "866734484103757845"
	log.Println(roleA, roleB)
	err = discord.GuildMemberRoleAdd(config.Guild, config.User, roleA)
	if err != nil {
		log.Println(err)
	}
}
