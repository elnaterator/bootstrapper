package input

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elnaterator/bootstrapper/pkg/color"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(color.Secondary.Hex))
)

func Choice(title string, options []string) string {

	fmt.Println(color.Term(title, color.Primary))

	items := mapToItems(options)

	l := list.New(items, itemDelegate{}, len(items)+1, 10)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	m := model{list: l}

	program := tea.NewProgram(m)
	out, err := program.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	choice := out.(model).choice
	fmt.Println(choice)
	return choice

}

func mapToItems(strs []string) []list.Item {
	items := make([]list.Item, len(strs))
	for i, s := range strs {
		items[i] = item(s)
	}
	return items
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	str := string(listItem.(item))
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" || m.quitting {
		return ""
	}
	return m.list.View()
}
