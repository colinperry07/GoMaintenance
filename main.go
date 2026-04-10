package main

import (
	"github.com/alecthomas/kong"
)

// to add:
// generate (gen) cmd
// view (view) cmd
// database management
// hashing passwords

var CLI struct {
	Stor struct {
		Location string `arg help:"Location of which these credentials belong."`
		Username string `arg help:"Your username, whether or not it's an email."`
		Password string `arg help:"Don't worry, this will be hashed later."`
	} `cmd help:"Stores your credential info into the db."`
}

func main() {
	kong.Parse(&CLI)
}
