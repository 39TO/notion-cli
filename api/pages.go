package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/39TO/notion-cli/types"
)

func GetPage(pageID string) (*types.Page, error) {
	url := fmt.Sprintf("%s/pages/%s", BaseURL, pageID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+NotionToken)
	req.Header.Set("Notion-Version", NotionVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var page types.Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func CreatePage(currentPageID, pageName string) error {
	url := fmt.Sprintf("%s/pages", BaseURL)
	payload := fmt.Sprintf(`{
		"parent": { "page_id": "%s" },
		"properties": {
			"title": [{ "text": { "content": "%s" } }]
		}
	}`, currentPageID, pageName)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+NotionToken)
	req.Header.Set("Notion-Version", NotionVersion)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("Created new page: %s\n", pageName)
	} else {
		return fmt.Errorf("error creating page: %s", resp.Status)
	}

	return nil
}
