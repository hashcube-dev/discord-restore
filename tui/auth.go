package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/spinner"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/hashcube-dev/discord-restore/discord"
)

type authModel struct{
	input textinput.Model
	spinner spinner.Model
	loading bool
	hidden bool
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

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return authModel{
		input: ti,
		spinner: sp,
		loading: false,
		hidden: true,
		width: 35,
		height: 20,
	}
}

func (m authModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  var cmds []tea.Cmd
	inhibitInput := false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.err = nil
			m.loading = true
			return m, authUser(m.input.Value())
		case "ctrl+h":
			m.hidden = !m.hidden
			if m.hidden {
				m.input.EchoMode = textinput.EchoPassword
			} else {
				m.input.EchoMode = textinput.EchoNormal
			}
			inhibitInput = true
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case AuthFail:
		m.err = msg.err
		m.loading = false
	case AuthPass:
		m.loading = false
	}
	if !m.loading {
		if !inhibitInput {
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		}
	} else {
		m.spinner, cmd = m.spinner.Update(m.spinner.Tick())
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m authModel) View() tea.View {
	var v tea.View
	var err string
	var loading string
	if m.err != nil {
		err = Error.Render(m.err.Error())
	}
	if m.loading {
		loading = Accented.Render(m.spinner.View(), "Loading")
	}
	v.Content = Default.
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(Default.Border(lipgloss.RoundedBorder(), true).Padding(2).
			Render(lipgloss.JoinVertical(lipgloss.Center,
				Accented.Render("Enter Discord Token"),
				Accented.Border(lipgloss.RoundedBorder()).Render(m.input.View()),
				loading,
				err,
			),
		),
	)

	return v
}

// Commands for Passing/Failing Auth
type AuthPass struct{}
type AuthFail struct{ err error }

func authUser(token string) tea.Cmd {
	return func() tea.Msg {
		var err error
		dc, err = discord.New(token)
		if err != nil {
			return AuthFail{err}
		}
		return AuthPass{}
	}
}
