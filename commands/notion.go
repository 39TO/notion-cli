package commands

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/39TO/notion-cli/utils"
)

func NotionOpen() {
	if utils.CurrentPageID == "" {
		fmt.Println("Error: No current page selected. Use 'cd' to navigate to a page first.")
		return
	}

	url := fmt.Sprintf("https://www.notion.so/%s", strings.ReplaceAll(utils.CurrentPageID, "-", ""))
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error opening browser: %v\n", err)
	} else {
		fmt.Printf("Opened current page in browser: %s\n", url)
	}
}
