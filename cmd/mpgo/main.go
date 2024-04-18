package main

import (
	"fmt"
	"mpgo/pkg/utils"
	"os"
	"path/filepath"
  "mpgo/pkg/music"
	tea "github.com/charmbracelet/bubbletea"
)



var (
  musicArray []string
)

type model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
}


func initialModel() model {
	return model{
		choices:  musicArray,
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
                    
          music.PlayMusic(musicArray[m.cursor])

        }
    }

    return m, nil
}


func (m model) View() string {
    s := "Select music\n\n"

    for i, choice := range m.choices {

        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }


        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }

    s += "\nPress q to quit.\n"

    return s
}


func main() {
  

  musicPath := utils.GetMusicDir()
  musicDir, _ := os.ReadDir(musicPath)


  for _, entry := range musicDir {
    currentFilePath := filepath.Join(musicPath, entry.Name())

    if entry.IsDir() {
      continue
    }
    
    if filepath.Ext(currentFilePath) != ".mp3" {
      continue
    }

    musicArray = append(musicArray, entry.Name())
  }


  p := tea.NewProgram(initialModel(), tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Alas, there's been an error: %v", err)
    os.Exit(1)
  }
}
