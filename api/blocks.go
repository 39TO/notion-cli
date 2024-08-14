package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/39TO/notion-cli/types"
)

func GetBlocks(pageID string) ([]types.Block, error) {
	url := fmt.Sprintf("%s/blocks/%s/children", BaseURL, pageID)
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

	var blocksResp types.BlocksResponse
	err = json.Unmarshal(body, &blocksResp)
	if err != nil {
		return nil, err
	}

	return blocksResp.Results, nil
}
