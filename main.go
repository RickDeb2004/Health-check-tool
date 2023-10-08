package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	
	"os"
	
)

func main() {
	app := &cli.App{
		Name:  "Health-checker",
		Usage: "A tiny tool that checks whether websites are running or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
			
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}

			
			domains := c.Args().Slice()

			if len(domains) == 0 {
				fmt.Println("Please provide at least one domain name as an argument.")
				return nil
			}

			
			resultChan := make(chan string)

			
			for _, domain := range domains {
				go func(domain string) {
					status := Check(domain, port)
					resultChan <- status
				}(domain)
			}

			
			for range domains {
				status := <-resultChan
				fmt.Println(status)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
