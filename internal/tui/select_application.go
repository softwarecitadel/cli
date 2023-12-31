package tui

import (
	"fmt"
	"os"

	"citadel/internal/api"

	tea "github.com/charmbracelet/bubbletea"
)

func newChooseApplicationPromptModel(projectSlug string) SelectModel {
	applications, err := api.RetrieveApplications(projectSlug)
	if err != nil {
		fmt.Println("Failed to retrieve applications")
		os.Exit(1)
	}

	choices := []SelectChoice{}
	for _, application := range applications {
		choices = append(choices, SelectChoice{
			Name: application.Name,
			ID:   application.ID,
			Slug: application.Slug,
		})
	}
	choices = append(choices, SelectChoice{
		Name: "Create a new application",
		ID:   "",
		Slug: "",
	})

	return NewSelectModel("Which application would you like to deploy to?", choices)
}

func SelectApplication(projectSlug string) string {
	m := newChooseApplicationPromptModel(projectSlug)
	res, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return res.(SelectModel).Choice.Slug
}
