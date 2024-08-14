package main

import (
	"wonderooo/spaj/v2/pkg/persist"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Add struct {
		Address string `arg:"" name:"address" help:"Address to spy." type:"address"`
	} `cmd:"" help:"Add address to spy."`
	Wipe struct {} `cmd:"" help:"Wipe all spies."`
}

func main() {
	db, err := persist.NewPlainTextDb("db-data")
	if err != nil {
		panic(err)
	}

	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "add <address>":
		db.Save(CLI.Add.Address)
	case "wipe":
		db.Wipe()
	default:
		panic(ctx.Command())
	}
}