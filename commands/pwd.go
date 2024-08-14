package commands

import (
	"fmt"

	"github.com/39TO/notion-cli/utils"
)

func Pwd() {
	fmt.Printf("Current page: %s\n", utils.CurrentPageName)
}
