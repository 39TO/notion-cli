package commands

import (
	"fmt"
	"strings"

	"github.com/39TO/notion-cli/api"
	"github.com/39TO/notion-cli/utils"
)

func Cd(pageName string) {
	if pageName == ".." {
		fmt.Println("Going up one level is not implemented yet.")
		return
	}

	blocks, err := api.GetBlocks(utils.CurrentPageID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, block := range blocks {
		if block.Type == "child_page" && strings.EqualFold(block.ChildPage.Title, pageName) {
			utils.CurrentPageID = block.ID
			utils.UpdateCurrentPageName()
			fmt.Printf("Moved to page: %s\n", utils.CurrentPageName)
			return
		}
	}

	fmt.Printf("Error: Page '%s' not found.\n", pageName)
}
