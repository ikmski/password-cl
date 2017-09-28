package main

import (
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func verifyPassword(c *cli.Context) error {

	return nil
}

func generatePassword(c *cli.Context) error {

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "password"
	app.Usage = "Password"
	app.Description = "Validate the password based on the policy. And generate a password"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:    "verify",
			Aliases: []string{"v"},
			Usage:   "verify password",
			Action:  verifyPassword,
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate password",
			Action:  generatePassword,
		},
	}

	app.RunAndExitOnError()
}
