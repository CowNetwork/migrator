package main

import (
	"log"
	"net/url"
	"os"
	"path"

	"github.com/amacneil/dbmate/pkg/dbmate"
)

func main() {
	postgresHost := os.Getenv("MIGRATOR_DB_ADDR")
	postgresUser := os.Getenv("MIGRATOR_DB_USER")
	postgresPassword := os.Getenv("MIGRATOR_DB_PASS")

	wrkdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mate := dbmate.New(&url.URL{
		Scheme:   "postgresql",
		Host:     postgresHost,
		User:     url.UserPassword(postgresUser, postgresPassword),
		Path:     "/" + os.Getenv("MIGRATOR_DB"),
		RawQuery: "&sslmode=" + os.Getenv("MIGRATOR_SSL_MODE"),
	})

	mate.MigrationsDir = path.Join(wrkdir, "db", "migrations")

	if err := mate.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := mate.Migrate(); err != nil {
		log.Fatal(err)
	}
}
