package main

import (
	"github.com/kyeongweonpark/papercoin/cli"
	"github.com/kyeongweonpark/papercoin/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
