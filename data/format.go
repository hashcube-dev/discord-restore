package data

import (
	discord "github.com/bwmarrin/discordgo"
)

type saveData struct{
	GuildID string `json:"guild"`
	GuildName string `json:"name"`
	ChannelMappings map[string]string `json:"channel_map"`
	UserMappings map[string]string `json:"name_map"`
}

func BuildSave(guild, name string, channels []*discord.Channel, members []*discord.Member) saveData {
	var save saveData

	save.GuildID = guild
	save.GuildName = name

	for _, channel := range channels {
		save.ChannelMappings[channel.ID] = channel.Name
	}
	for _, member := range members {
		save.UserMappings[member.User.ID] = member.Nick
	}
	
	return save
}
