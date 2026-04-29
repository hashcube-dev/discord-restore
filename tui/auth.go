package tui

import (
	"time"
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/spinner"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/hashcube-dev/discord-restore/discord"
)

type authModel struct{
	input textinput.Model
	spinner spinner.Model
	err error
	width int
	height int
}

func AuthModel() authModel {
	ti := textinput.New()
	ti.Focus()
	ti.EchoMode = textinput.EchoPassword

	return authModel{input: ti}
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
				disc := make(chan discord.DiscordController)
				err := make(chan error)
				defer close(disc)
				defer close(err)
				ticker := time.NewTicker(500 * time.Millisecond)

				go func(dcCh chan discord.DiscordController, errCh chan error) {
					dc, err := discord.New(m.input.Value())
					dcCh <- dc
					errCh <- err
				}(disc, err)
				loading:
				for {
					select {
					case <-disc:
						break loading
					case <-ticker.C:
						m.spinner.Tick()
					}
				}
				dc = <-disc
				if <-err != nil {
					m.err = <-err
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
	v.Content = Default.
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(Default.Border(lipgloss.RoundedBorder(), true).Padding(2).
			Render(lipgloss.JoinVertical(lipgloss.Center,
				Accented.Render("Enter Discord Token"),
				Default.Render(m.input.View()),
				Error.Render(m.err.Error()),
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
