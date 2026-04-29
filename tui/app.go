package tui

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	ctp "github.com/catppuccin/go"
	log "charm.land/log/v2"
	"github.com/hashcube-dev/discord-restore/discord"
)

var dc discord.DiscordController

type Page uint

const (
  Auth Page = iota
	Save
	Load
	Restore
	Settings
	Cmd
)

var Flavor ctp.Flavor
var Accent ctp.Color

var Error lipgloss.Style
var Warning lipgloss.Style
var Success lipgloss.Style
var Accented lipgloss.Style
var Default lipgloss.Style

type mainModel struct{
	CurrentlyViewing Page

	Auth tea.Model
	Save tea.Model
}

func StartTea() {

	Flavor = ctp.Mocha
	Accent = Flavor.Pink()

	Default = lipgloss.NewStyle().
		MarginBackground(lipgloss.Color(Flavor.Base().Hex)).
		Background(lipgloss.Color(Flavor.Base().Hex)).
		Foreground(lipgloss.Color(Flavor.Text().Hex)).
		BorderForeground(lipgloss.Color(Accent.Hex)).
		BorderBackground(lipgloss.Color(Flavor.Base().Hex))
	Error = Default.
		Bold(true).
		Foreground(lipgloss.Color(Flavor.Red().Hex))
	Warning = Default.
		Bold(true).
		Foreground(lipgloss.Color(Flavor.Yellow().Hex))
	Success = Default.
		Bold(true).
		Foreground(lipgloss.Color(Flavor.Green().Hex))
	Accented = Default.
		Bold(true).
		Foreground(lipgloss.Color(Accent.Hex))
	p := tea.NewProgram(NewMainModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func NewMainModel() tea.Model {
	return mainModel{
		CurrentlyViewing: Auth,

		Auth: AuthModel(),
  }
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msgc := msg.(type) {
	case tea.KeyMsg:
		switch msgc.String() {
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		default:
			switch m.CurrentlyViewing {
			case Auth:
				m.Auth, cmd = m.Auth.Update(msg)
			}
		}
	case tea.WindowSizeMsg:
		switch m.CurrentlyViewing {
		case Auth:
			m.Auth, cmd = m.Auth.Update(msg)
		}
	case AuthPass:
		m.CurrentlyViewing = Save
		m.Auth, cmd = m.Auth.Update(msg)
	default: 
		switch m.CurrentlyViewing {
		case Auth:
			m.Auth, cmd = m.Auth.Update(msg)
		}
	}

	return m, tea.Batch(cmd)
}

func (m mainModel) View() tea.View {
	var v tea.View
	switch m.CurrentlyViewing {
	case Auth:
		v = m.Auth.View()
	}
	v.AltScreen = true
	return v
}
