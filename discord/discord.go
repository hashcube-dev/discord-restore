package discord

import (
	discord "github.com/bwmarrin/discordgo"
)

type DiscordController struct{
	session *discord.Session
}

func New(token string) (DiscordController, error) {
	var d DiscordController
	s, err := discord.New(token)
	if err != nil {
		return d, err
	}
	d.session = s
	_, err = d.session.User("@me")
	return d, err
}

func (c *DiscordController) GetChannels(guild string) ([]*discord.Channel, error) {
	channels, err := c.session.GuildChannels(guild)
	return channels, err
}

func (c *DiscordController) GetUsers(guild string) ([]*discord.Member, error) {
	members, err := c.session.GuildMembers(guild, "0", 1000)
	return members, err
}

func (c *DiscordController) GetUserGuilds() ([]*discord.UserGuild, error) {
	ug, err := c.session.UserGuilds(200, "", "", false)
	return ug, err
} 

func (c *DiscordController) ChangeNick(guild, user, nickname string) error {
	err := c.session.GuildMemberNickname(guild, user, nickname)
	return err
}

func (c *DiscordController) EditChannelName(guild, channel, name string) (*discord.Channel, error) {
	var chEdit discord.ChannelEdit
	chEdit.Name = name
	ch, err := c.session.ChannelEdit(channel, &chEdit)
	return ch, err
}
