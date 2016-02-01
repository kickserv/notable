package main

import (
	notable "github.com/joshualcoffee/notable"
	"os"
)

func main() {
	if len(notable.Notes()) > 0 {
		notable.SendEmail()
		notable.Reset()
	}
}
