package main

import (
	"log"
	"os"

	. "github.com/amar-jay/nat_wsl/pkg/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	ipTypes := []string{"v4tov4", "v4tov6", "v6tov4", "v6tov6"}

	app := &cli.App{
		Name:  "nat_wsl",
		Usage: "A tool to forward ports from WSL to Windows",
		Commands: []*cli.Command{
			{
				Name:    "configfile",
				Aliases: []string{"config"},
				Usage:   "Load configuration from env variable NATWSL_CONFIG_PATH",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "config_path",
						Category: "Custom: ",
						Value:    os.Getenv("NATWSL_CONFIG_PATH"),
						Usage:    "Load configuration from `FILE`",
						Action: func(ctx *cli.Context, s string) error {
							if s == "" {
								// use env variable
								return errors.Errorf("`FILE` variable not set")
							}

							// check if file path exists
							if _, err := os.Stat(s); os.IsNotExist(err) {
								return errors.Errorf("provided config path does not exist")
							}

							return ctx.Set("config_path", s)
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					config := Config{}
					file_path := cCtx.String("config_path")
					// check if file path exists
					if _, err := os.Stat(file_path); os.IsNotExist(err) {
						log.Fatal("NATWSL_CONFIG_PATH does not exist")
					}

					// take file content and yaml unmarshal
					content, err := os.ReadFile(file_path)
					if err != nil {
						log.Fatalf("Can't read file: %v", err)
					}

					yaml.Unmarshal(content, &config)
					config.SetDefaults()
					log.Printf("config: %+v", config)

					return nil
				},
			},
			{
				Name:  "portproxy",
				Usage: "Forward ports from WSL to Windows",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "protocol",
						Value:    "tcp",
						Category: "Custom: ",
						Usage:    "type of protocol connection",
						Action: func(ctx *cli.Context, s string) error {
							if s == "tcp" || s == "udp" {
								return nil
							}
							return errors.Errorf("Invalid protocol")
						},
					},
					&cli.StringFlag{
						Name:     "type",
						Value:    "v4tov4",
						Usage:    "type of ip connection",
						Category: "Custom: ",
						Action: func(ctx *cli.Context, s string) error {
							for _, ipType := range ipTypes {
								if s == ipType {
									return nil
								}
							}
							return errors.Errorf("Invalid ip type")
						},
					},

					&cli.StringFlag{
						Name:     "connecthost",
						Required: true,
						Aliases:  []string{"ch"},
						Usage:    "Remote `REMOTE_IP` to listen to",
					},
					&cli.StringFlag{
						Name:     "connectport",
						Required: true,
						Aliases:  []string{"cp"},
						Usage:    "Listen to `REMOTE_PORT` of remote host",
					},
					&cli.StringFlag{
						Name:     "listenhost",
						Required: true,
						Aliases:  []string{"lh"},
						Usage:    "WSL `WSL_IP` to connect to",
					},
					&cli.StringFlag{
						Name:     "listenport",
						Required: true,
						Aliases:  []string{"lp"},
						Usage:    "Connect to `WSL_PORT` in WSL",
					},
				},
				Action: func(cCtx *cli.Context) error {
					portproxy := Forwarding{}
					portproxy.Protocol = cCtx.String("protocol")
					portproxy.Type = cCtx.String("type")
					portproxy.Wsl.Listenport = cCtx.Int("listenport")
					portproxy.Remote.Connectport = cCtx.Int("connectport")
					portproxy.Wsl.Listenhost = cCtx.String("listenhost")
					portproxy.Remote.Connectip = cCtx.String("connecthost")

					conf := Config{
						"portproxy": portproxy,
					}
					log.Printf("config: %+v", conf)

					return nil
				},
			},
		},

		Action: func(cCtx *cli.Context) error {
			println(cCtx.App.HelpName, ":", cCtx.App.Usage)
			println("Type 'nat_wsl --help' to see available commands")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
