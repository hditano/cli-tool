package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/hditano/cli-tool/utils"
)

type actor struct {
	name    string
	age     int16
	species string
}

type animals interface {
	eat()
	noise() string
}

func (a actor) eat() {
	fmt.Println(a)
}

func (a actor) noise() string {
	return fmt.Sprintf("Data is %s %v %s: ", a.name, a.age, a.species)
}

func animalsActions(a animals) {
	fmt.Printf("%s", a.noise())
	a.eat()
}

func actorAction(ctx context.Context, cmd *cli.Command) error {
	p1 := actor{"Hernan", 42, "Human"}

	animalsActions(p1)
	return nil
}

var rootCmd *cli.Command

func init() {
	rootCmd = &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "boom",
				Usage: "Main group of explosive commands",
				Commands: []*cli.Command{
					{
						Name:    "subcommand1",
						Usage:   "First explosive action",
						Aliases: []string{"sub1"},
						Action:  utils.BoomAction,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "lang",
								Value:   "English",
								Usage:   "Idioma para el saludo",
								Aliases: []string{"l"},
								Sources: cli.EnvVars("APP_LANG"),
							},
						},
					},
					{
						Name:    "subcommand2",
						Usage:   "Second explosive action",
						Aliases: []string{"sub2"},
						Action:  utils.ListFiles,
					},
				},
			},
			{
				Name:   "actor",
				Usage:  "Testing interfaces",
				Action: actorAction,
			},
			{
				Name:  "database",
				Usage: "testing sqlite database",
				Commands: []*cli.Command{
					{
						Name:   "createDB",
						Usage:  "Create Sqlite3 DB",
						Action: utils.CreateDB,
					},
					{
						Name:   "CheckConnection",
						Usage:  "Check Sqlite3 DB Connection",
						Action: utils.CheckDBConnection,
					},
				},
			},
		},
	}
}

func Execute() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
