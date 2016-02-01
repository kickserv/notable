package notable

import (
	"log"
	"net/mail"
	postmark "github.com/hjr265/postmark.go/postmark"
	"os"
	"strings"
)

func sendViaPostmark(subject string, body string) {
	client := postmark.Client{
		ApiKey: os.Getenv("POSTMARK_API_KEY"),
		Secure: true,
	}
	html := strings.NewReader(body)

	res, err := client.Send(&postmark.Message{
		From: &mail.Address{
			Name:    os.Getenv("FROM_NAME"),
			Address: os.Getenv("FROM_EMAIL"),
		},
		To: []*mail.Address{
			{
				Name:    os.Getenv("TO_NAME"),
				Address: os.Getenv("TO_EMAIL"),
			},
		},
		Subject:  subject,
		HtmlBody: html,
	})

	if err != nil {
		log.Print("postmark errors")
		log.Print(err)
	}

	if res != nil {
		log.Print("postmark success")
		log.Print(res)
	}
}
