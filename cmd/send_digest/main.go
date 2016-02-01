package main

import (
		notable "github.com/harvesthq/notable"
)

func main() {
	if len(notable.Notes()) > 0 {
		notable.SendEmail()
		notable.Reset()
	}
}
