package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cheveuxdelin/cipher/caesar"
	"github.com/cheveuxdelin/cipher/morse"
	"github.com/urfave/cli/v2"
)

func encodeMorse(c *cli.Context) error {

	if c.Args().Len() == 0 {
		return errors.New("no string provided")
	}

	value, error := morse.Encode(c.Args().First())

	if error == nil {
		fmt.Println(value)
	} else {
		return error
	}

	return nil
}

func decodeMorse(c *cli.Context) error {

	if c.Args().Len() == 0 {
		return errors.New("no text provided")
	}

	value, error := morse.Decode(c.Args().First())

	if error == nil {
		fmt.Println(value)
	} else {
		return error
	}

	return nil
}

func encodeCaesar(c *cli.Context) error {

	if c.NArg() == 0 {
		return errors.New("no message provided")
	} else if c.NArg() == 1 {
		return errors.New("incomplete arguments")
	}

	var message string = c.Args().First()
	n, error := strconv.Atoi(c.Args().Get(1))

	if error != nil {
		return error
	}

	fmt.Println(caesar.Encode(message, byte(n), c.Bool("onlyletters")))

	return nil
}

func decodeCaesar(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("no message provided")
	} else if c.NArg() == 1 {
		return errors.New("incomplete arguments")
	}

	var message string = c.Args().First()
	n, error := strconv.Atoi(c.Args().Get(1))

	if error != nil {
		return error
	}

	fmt.Println(caesar.Decode(message, byte(n), c.Bool("onlyletters")))

	return nil
}

func main() {

	app := &cli.App{
		Name:  "cipher",
		Usage: "CLI App to encode and decode in different encoding systems :).",
		Commands: []*cli.Command{
			{
				Name:    "encode",
				Aliases: []string{"e"},
				Usage:   "Encodes incoming message",
				Subcommands: []*cli.Command{
					// morse
					{
						Name:    "morse",
						Aliases: []string{"m"},
						Action:  encodeMorse,
					},
					// caesar
					{
						Name:    "caesar",
						Aliases: []string{"c"},
						Action:  encodeCaesar,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "onlyletters",
								Aliases: []string{"o"},
								Usage:   "encode only letters from `MESSAGE`",
							},
						},
					},
				},
			},
			{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "Decodes incoming message",
				Subcommands: []*cli.Command{
					// morse
					{
						Name:    "morse",
						Aliases: []string{},
						Action:  decodeMorse,
					},
					// morse
					{
						Name:    "caesar",
						Aliases: []string{},
						Action:  decodeCaesar,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
