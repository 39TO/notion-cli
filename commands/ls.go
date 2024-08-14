package commands

import (
	"fmt"

	"github.com/39TO/notion-cli/api"
	"github.com/39TO/notion-cli/utils"
)

func Ls() {
	if utils.CurrentPageID == "" {
		fmt.Println("Error: No current page selected. Use 'cd' to navigate to a page first.")
		return
	}

	blocks, err := api.GetBlocks(utils.CurrentPageID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, block := range blocks {
		if block.Type == "child_page" {
			fmt.Printf("%s (Page)\n", block.ChildPage.Title)
		} else if block.Type == "paragraph" && len(block.Paragraph.RichText) > 0 {
			fmt.Printf("%s (Paragraph)\n", block.Paragraph.RichText[0].PlainText)
		}
	}
}
