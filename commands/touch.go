package commands

import (
	"fmt"

	"github.com/39TO/notion-cli/utils"

	"github.com/39TO/notion-cli/api"
)

func Touch(pageName string) {
	if utils.CurrentPageID == "" {
		fmt.Println("Error: No current page selected. Use 'cd' to navigate to a page first.")
		return
	}

	err := api.CreatePage(utils.CurrentPageID, pageName)
	if err != nil {
		fmt.Printf("Error creating page: %v\n", err)
	}
}
