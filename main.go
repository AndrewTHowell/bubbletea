package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(newChoiceModel(
		"Apples",
		"Bananas",
		"Oranges",
	))
	if _, err := program.Run(); err != nil {
		panic(any(err))
	}
}

func newChoiceModel(choices ...string) *choiceModel {
	return &choiceModel{
		choices:  choices,
		cursor:   0,
		selected: make(map[int]bool, len(choices)),
	}
}

type choiceModel struct {
	choices  []string
	cursor   int
	selected map[int]bool
}

func (m choiceModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m choiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}
		}
	}
	return m, nil
}

func (m choiceModel) View() string {
	var str string

	str += fmt.Sprintln("*** Hello there ***")
	str += fmt.Sprintln()
	for i, choice := range m.choices {

		// Render cursor
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Render selection
		selected := "  "
		if m.selected[i] {
			selected = "x"
		}

		str += fmt.Sprintln(fmt.Sprintf("%s [%s] %s", cursor, selected, choice))
	}
	str += fmt.Sprintln()
	str += fmt.Sprintln("Press q to quit")
	return str
}
