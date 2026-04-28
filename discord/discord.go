package discord

import (
	discord "github.com/bwmarrin/discordgo"
)

type DiscordController struct{
	session *discord.Session
	guild string
}

func InitSession(token, guild string) (DiscordController, error) {
	var d DiscordController
	s, err := discord.New(token)
	if err != nil {
		return d, err
	}
	d.session = s
	return d, nil
}

func (c *DiscordController) GetChannels() ([]*discord.Channel, error) {
	channels, err := c.session.GuildChannels(c.guild)
	return channels, err
}

func (c *DiscordController) GetUsers() ([]*discord.Member, error) {
	members, err := c.session.GuildMembers(c.guild, "0", 1000)
	return members, err
}

func (c *DiscordController) ChangeNick(user, nickname string) error {
	err := c.session.GuildMemberNickname(c.guild, user, nickname)
	return err
}

func (c *DiscordController) EditChannelName(channel, name string) error {
	// TODO: Figure out how to edit the channel name
	return nil
}
