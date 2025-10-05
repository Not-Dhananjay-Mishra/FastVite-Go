package utils

import (
	"fmt"
	"govite/process"

	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	selectMode mode = iota
	inputMode
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	mode     mode
	input    string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case selectMode:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "enter":
				// If user presses enter on a selection, go to input mode
				m.mode = inputMode
			case " ":
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}

	case inputMode:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyRunes:
				m.input += string(msg.Runes)
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyEnter:
				sus := fmt.Sprintf("Name of the Web Application: %s", m.input)
				fmt.Println(sus)
				process.GoAndJS(m.input)
				return m, tea.Quit
			case tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.mode == selectMode {
		s := "Tech-Stack preferred (use space to select, enter to continue)\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}
		s += "\nPress q to quit.\n"
		return s
	}

	if m.mode == inputMode {
		sus := fmt.Sprintf("%v\n\nName of the Web Application: %s", m.choices[0], m.input)
		return sus
	}

	return ""
}

func (m model) Init() tea.Cmd {

	return nil
}

func InitialModel() model {
	return model{
		choices:  []string{"Vite + Go (Javascript)", "Vite + Go (Typescript)"},
		selected: make(map[int]struct{}),
		mode:     selectMode,
		input:    "",
	}
}
