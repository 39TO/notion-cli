package utils

import (
	"fmt"

	"github.com/39TO/notion-cli/api"
	"github.com/39TO/notion-cli/config"
)

var (
	CurrentPageID   string
	CurrentPageName string = "/"
)

func InitState() {
	CurrentPageID = config.GetNotionRootPageID()
}

func UpdateCurrentPageName() {
	page, err := api.GetPage(CurrentPageID)
	if err != nil {
		fmt.Printf("Error fetching page info: %v\n", err)
		CurrentPageName = "Unknown"
		return
	}
	if len(page.Properties.Title.Title) > 0 {
		CurrentPageName = page.Properties.Title.Title[0].PlainText
	} else {
		CurrentPageName = "Untitled"
	}
}
