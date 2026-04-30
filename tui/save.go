package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/list"
	discord "github.com/bwmarrin/discordgo"
	"charm.land/log/v2"

	"strings"
	"fmt"
	"io"
)

type saveModel struct{
	gdList list.Model
	chList list.Model
}

type channelItem struct{
	ch *discord.Channel
}

type guildItem struct{
	ch *discord.UserGuild
}

func (i channelItem) FilterValue() string { return i.ch.Name }
func (i   guildItem) FilterValue() string { return i.ch.Name }

type channelDelegate struct {
	cursor rune
	selected rune
}

func (d channelDelegate) Height() int                             { return 1 }
func (d channelDelegate) Spacing() int                            { return 0 }
func (d channelDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d channelDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(channelItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("   %s", i.ch.Name)

	fn := func(s ...string) string {
		return Default.Render("  " + strings.Join(s, ""))
	}
	if index == m.Index() {
		fn = func(s ...string) string {
			return Accented.Render(" " + string(d.cursor) + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type guildDelegate struct {
	cursor rune
}

func (d guildDelegate) Height() int                             { return 1 }
func (d guildDelegate) Spacing() int                            { return 0 }
func (d guildDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d guildDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(guildItem)
	if !ok {
		return
	}

	str := fmt.Sprintf(" %s", i.ch.Name)

	fn := func(s ...string) string {
		return Default.Render("  " + strings.Join(s, ""))
	}
	if index == m.Index() {
		fn = func(s ...string) string {
			return Accented.Render(" " + string(d.cursor) + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func DiscordChannelListToItemList(channels []*discord.Channel) []list.Item {
	var itemList []list.Item
	for _, ch := range channels {
		itemList = append(itemList, channelItem{ch})
	}
	return itemList
}

func DiscordUserGuildListToItemList(channels []*discord.UserGuild) []list.Item {
	var itemList []list.Item
	for _, ch := range channels {
		itemList = append(itemList, guildItem{ch})
	}
	return itemList
}

func SaveModel() saveModel {
	cd := channelDelegate{
		cursor: '>',
	}
	gd := guildDelegate{
		cursor: '>',
	}
	var chList []list.Item
	var gdList []list.Item

	guilds, err := dc.GetUserGuilds()
	if err != nil {
		log.Fatal(err)
	}
	gdList = DiscordUserGuildListToItemList(guilds)

	chli := list.New(chList, cd, 50, 50)
	chli.SetShowHelp(false)
	gdli := list.New(gdList, gd, 50, 50)
	gdli.SetShowHelp(false)
	return saveModel{
		chList: chli,
		gdList: gdli,
	}
}

func (m saveModel) Init() tea.Cmd {
	return nil
}

func (m saveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
	case tea.WindowSizeMsg:
		m.gdList.SetSize(msg.Width/2, msg.Height/2)
		m.chList.SetSize(msg.Width/2, msg.Height/2)
	}

	var cmd tea.Cmd
	m.gdList, cmd = m.gdList.Update(msg)
	m.chList, cmd = m.chList.Update(msg)
	return m, cmd
}

func (m saveModel) View() tea.View {
	var v tea.View
	v.Content = Default.Render(m.gdList.View())
	return v
}
