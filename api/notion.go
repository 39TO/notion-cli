package api

import (
	"github.com/39TO/notion-cli/config"
)

var (
	NotionToken   string
	NotionVersion string = "2022-06-28"
	BaseURL       string = "https://api.notion.com/v1"
)

func InitAPI() {
	NotionToken = config.GetNotionToken()
}
