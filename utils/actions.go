package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"

	_ "github.com/mattn/go-sqlite3"
)

func BoomAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.String("lang") == "English" {
		fmt.Println("Hello from English")
	}
	if cmd.String("lang") == "Spanish" {
		fmt.Println("Hola desde Spanish")
	}
	file, err := os.ReadFile("./go.sum")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(file)
	return nil

}

func CreateDB(ctx context.Context, cmd *cli.Command) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func CheckDBConnection(ctx context.Context, cmd *cli.Command) error {
	db, err := sql.Open("sqlite3", "file:./foo.db?mode=ro&_busy_timeoput=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	fmt.Println("DB Connection succesfully")

	return nil
}

func ListFiles(ctx context.Context, cmd *cli.Command) error {
	req, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req)
	return nil
}
