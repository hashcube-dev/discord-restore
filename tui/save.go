package tui

import (
	tea "charm.land/bubbletea/v2"
)

type saveModel struct{
	// TODO:
}

func SaveModel() saveModel {
	return saveModel{
		// TODO: 
	}
}

func (m saveModel) Init() tea.Cmd {
	return nil
}

func (m saveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m saveModel) View() tea.View {
	var v tea.View
	return v
}
