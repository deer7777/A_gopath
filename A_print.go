package A_gopath

import (
	"fmt"
	PrintNow "github.com/deer7777/B"
	"github.com/deer7777/A_gopath/utils"
)

func A_print(node *utils.node) {
	fmt.Println("version_A_gopath")
	PrintNow.PrintNow()
	fmt.Println(node)
}
