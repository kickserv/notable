package main

import (
	notable "notable/Godeps/_workspace/src/github.com/joshualcoffee/notable"
)

func main() {
	if len(notable.Notes()) > 0 {
		notable.SendEmail()
		notable.Reset()
	}
}
