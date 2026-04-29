package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/textinput"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/hashcube-dev/discord-restore/discord"
)

type authModel struct{
	input textinput.Model
	err error
	width int
	height int
}

func AuthModel() authModel {
	ti := textinput.New()
	ti.Focus()
	ti.Prompt = ""
	ti.EchoMode = textinput.EchoPassword
	ti.SetWidth(70)

	return authModel{
		input: ti,
		width: 35,
		height: 20,
	}
}

func (m authModel) Init() tea.Cmd {
	return nil
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
			case tea.KeyEnter:
				var err error
				dc, err = discord.New(m.input.Value())
				if err != nil {
					m.err = err
					return m, authFail()
				}
				return m, authPass()
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m authModel) View() tea.View {
	var v tea.View
	var err string
	if m.err != nil {
		err = Error.Render(m.err.Error())
	}
	v.Content = Default.
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(Default.Border(lipgloss.RoundedBorder(), true).Padding(2).
			Render(lipgloss.JoinVertical(lipgloss.Center,
				Accented.Render("Enter Discord Token"),
				Accented.Border(lipgloss.RoundedBorder()).Render(m.input.View()),
				err,
			),
		),
	)

	return v
}

// Commands for Passing/Failing Auth
type AuthPass struct{}

func authPass() tea.Cmd {
	return func() tea.Msg {
		return AuthPass{}
	}
}

type AuthFail struct{}

func authFail() tea.Cmd {
	return func() tea.Msg {
		return AuthFail{}
	}
}
