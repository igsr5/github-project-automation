package command

import (
	"encoding/json"
	"fmt"
	"os/exec"

	variables "github.com/sora-ichigo/github-project-automation"
	usecase "github.com/sora-ichigo/github-project-automation/usecase"
)

type Item struct {
	ID string `json:"id"`
}

type projectV2SetterImpl struct {
}

// NewProjectV2Setter is a factory method to create a new instance of ProjectV2Setter.
func NewProjectV2Setter() usecase.ProjectV2Setter {
	return &projectV2SetterImpl{}
}

// Set sets project items.
func (s *projectV2SetterImpl) Set(categoryID string, statusID string, projectItems []usecase.ProjectItem) error {
	fmt.Printf("projectItems: %v\n", projectItems)
	// 1. Add project items
	items := make([]Item, 0, len(projectItems))
	for _, item := range projectItems {
		item, err := addProjectItem(item.URL)
		if err != nil {
			return fmt.Errorf("failed to add project item: %w", err)
		}
		items = append(items, *item)
	}

	fmt.Printf("items: %v\n", items)
	// 2. Move project item status to specified status
	for _, id := range items {
		err := moveStatusProjectItem(id, statusID)
		if err != nil {
			return fmt.Errorf("failed to move status project item: %w", err)
		}
	}

	// 3. Move project item category to specified category
	for _, id := range items {
		err := moveCategoryProjectItem(id, categoryID)
		if err != nil {
			return fmt.Errorf("failed to move category project item: %w", err)
		}
	}

	return nil
}

func addProjectItem(url string) (*Item, error) {
	cmd := exec.Command("gh", "project", "item-add", variables.PROJECT_NUMBER, "--owner", "wantedly", "--url", url, "--format", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to add project item: %w", err)
	}

	var item Item
	if err := json.Unmarshal(output, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal project item: %w", err)
	}

	return &item, nil
}

func moveStatusProjectItem(item Item, statusID string) error {
	cmd := exec.Command("gh", "project", "item-edit", "--id", item.ID, "--field-id", variables.STATUS_FIELD_ID, "--project-id", variables.PROJECT_ID, "--single-select-option-id", statusID)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to move status project item: %w", err)
	}

	return nil
}

func moveCategoryProjectItem(item Item, categoryID string) error {
	cmd := exec.Command("gh", "project", "item-edit", "--id", item.ID, "--field-id", variables.CATEGORY_FIELD_ID, "--project-id", variables.PROJECT_ID, "--single-select-option-id", categoryID)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to move category project item: %w", err)
	}

	return nil
}
