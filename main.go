package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/39TO/notion-cli/api"
	"github.com/39TO/notion-cli/commands"
	"github.com/39TO/notion-cli/utils"
)

func main() {
	api.InitAPI()
	utils.InitState()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", utils.CurrentPageName)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "exit":
			return
		case "ls":
			commands.Ls()
		case "cd":
			if len(args) < 2 {
				fmt.Println("Error: 'cd' command requires a page name.")
			} else {
				commands.Cd(args[1])
			}
		case "touch":
			if len(args) < 2 {
				fmt.Println("Error: 'touch' command requires a page name.")
			} else {
				commands.Touch(args[1])
			}
		case "notion":
			commands.NotionOpen()
		case "pwd":
			commands.Pwd()
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  ls               - List contents of the current page")
			fmt.Println("  cd <page_name>   - Change to a subpage")
			fmt.Println("  touch <page_name> - Create a new page")
			fmt.Println("  notion           - Open current page in browser")
			fmt.Println("  pwd              - Print current page name")
			fmt.Println("  exit             - Exit the program")
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	}
}
