package main

import (
	"crypto/tls"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Message struct {
	body []byte
}

func getMessages() {
	email := os.Getenv("NONAME_EMAIL")
	password := os.Getenv("NONAME_PASS")

	c := newClient()
	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(email, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select mailbox
	mbox, err := c.Select("carris", false)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Messages)

	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}

	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}
	if subject, err := header.Subject(); err == nil {
		log.Println("Subject:", subject)
	}

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error iterating mail part: ", err)
		}

		switch p.Header.(type) {
		case mail.TextHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			log.Printf("Got text: %v\n", string(b))
		}
	}
}

func newClient() *client.Client {
	host := os.Getenv("NONAME_IMAP_HOST")
	port := os.Getenv("NONAME_IMAP_PORT")

	url := fmt.Sprintf("%v:%v", host, port)
	log.Printf("Connecting to IMAP server %v...\n", url)

	// Connect to server
	config := tls.Config{InsecureSkipVerify: true}
	c, err := client.DialTLS(url, &config)
	if err != nil {
		log.Fatalln("Error connecting to IMAP server: ", err)
	}
	log.Println("Connected")

	return c
}
