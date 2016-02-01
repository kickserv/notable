package notable

import (
	"log"
	mandrill "notable/Godeps/_workspace/src/github.com/keighl/mandrill"
	"os"
)

func sendViaMandrill(subject string, body string) {
	client := mandrill.ClientWithKey(os.Getenv("MANDRILL_API_KEY"))

	message := &mandrill.Message{}
	message.AddRecipient(os.Getenv("TO_EMAIL"), os.Getenv("TO_NAME"), "to")
	message.FromEmail = os.Getenv("FROM_EMAIL")
	message.FromName = os.Getenv("FROM_NAME")
	message.Subject = subject
	message.HTML = body

	_, err := client.MessagesSend(message)
	if err != nil {
		log.Print(err)
	}
}
