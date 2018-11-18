package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"raqb.it/catbot/utils"
	"strings"
)

const ckBaseUrl = "https://storage.googleapis.com/ck-kitty-image/0x06012c8cf97bead5deae237070f9587f8e7a266d"

func Info(s *discordgo.Session, m *discordgo.MessageCreate,
	parts []string, globalEnv *GlobalEnv, cmdEnv *CommandEnv) (error) {

	if len(parts) < 2 {
		utils.ChannelMesageSendError(s, m.ChannelID, "Please specify a cat to get the info of!")
		return nil
	}

	catName := strings.Join(parts[1:], " ")

	cat, err := globalEnv.Db.GetCatByName(cmdEnv.User.ID, catName)

	if err != nil && cat == nil {
		utils.ChannelMesageSendError(s, m.ChannelID, fmt.Sprintf(
			"%s, you do not have a cat with that name!",
			m.Author.Mention(),
		))
		return nil
	}

	catStatus := "Home"

	if cat.Away {
		catStatus = "Away"
	}

	embed := &discordgo.MessageEmbed{
		Title: cat.Name,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("%s/%d.png", ckBaseUrl, cat.CryptoKittyID),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Pronoun",
				Value: cat.Pronoun,
			},
			{
				Name: "Status",
				Value: catStatus,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return nil
}
