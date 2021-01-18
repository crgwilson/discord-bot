package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

const usageText = `botdb: Perform database migrations

Usage:
    go run *.go [flags] <command>

Flags:
    -H                       Host address of the PostgreSQL database
    -p                       Host port of the PostgreSQL database
    -U                       Login user for the PostgreSQL database
    -P                       Login password for the PostgreSQL database
    -d                       The name of the database to connect to

Commands:
    init                     creates version info table in the database
    up                       runs all available migrations.
    up [target]              runs available migrations up to the target one.
    down                     reverts last migration.
    reset                    reverts all migrations.
    version                  prints current db version.
    set_version [version]    sets db version without running migrations.

`

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	dbHost := flag.String("H", "localhost", "Host address of the PostgreSQL database")
	dbPort := flag.Int("p", 5432, "Host port of the PostgreSQL database")
	dbUser := flag.String("U", "postgres", "Login user for the PostgreSQL database")
	dbPassword := flag.String("P", "", "Login password for the PostgreSQL database")
	dbName := flag.String("d", "postgres", "The name of the database to connect to")

	flag.Usage = usage

	flag.Parse()

	dbAddr := fmt.Sprintf("%s:%d", *dbHost, *dbPort)
	dbOptions := pg.Options{
		User:     *dbUser,
		Password: *dbPassword,
		Database: *dbName,
		Addr:     dbAddr,
	}

	db := pg.Connect(&dbOptions)

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		panic(err)
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}
