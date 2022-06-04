package main

import (
	"cnc/colours"
	"context"
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func tryLogin(term *term.Terminal) {
	term.SetPrompt("Username ")

	username, err := term.ReadLine()
	if err != nil {
		fmt.Print(err)
	}
	term.SetPrompt("Password ")

	password, err := term.ReadPassword("Password ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(username)
	fmt.Print(password)

	var HashedPassword string

	err = DB.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", username).Scan(&HashedPassword)

	if err != nil {
		fmt.Print()
		fmt.Print(err)

		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password)); err != nil {
		fmt.Print("INCOORECT")
		return
	}

}

func StartServer() {
	ssh.Handle(func(s ssh.Session) {
		term := term.NewTerminal(s, "")
		tryLogin(term)
		term.SetPrompt(fmt.Sprintf("[%s%s%s@%sCNC%s]$ ", colours.Blue, "ralph", colours.White, colours.Red, colours.Reset))
		for {
			command, err := term.ReadLine()

			if err != nil {
				break
			}

			log.Println(fmt.Sprintf("%s ran command \"%s\"", s.RemoteAddr(), command))

		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/Users/hh/.ssh/id_rsa")))
}
