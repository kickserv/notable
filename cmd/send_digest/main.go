package main

import (
	notable "github.com/joshualcoffee/notable"
)

func main() {
	if len(notable.Notes()) > 0 {
		notable.SendEmail()
		notable.Reset()
	}
}
