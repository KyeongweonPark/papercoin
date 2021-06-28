package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/semipia/papercoin/explorer"
	"github.com/semipia/papercoin/rest"
)

// import	"github.com/semipia/papercoin/rest"

// import "github.com/semipia/papercoin/explorer"

func usage() {
	fmt.Printf("Welcome to 페이퍼코인\n\n")
	fmt.Printf("use the following flags:\n\n")
	fmt.Printf("-port=4000:	 Set the PORT of the server\n")
	fmt.Printf("-mode=rest:  Choose between 'html' and 'rest' \n\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port fo the server")
	mode := flag.String("mode", "rest", "choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
		// rest API
	case "html":
		explorer.Start(*port)
		// HTML exp
	default:
		usage()
	}
	fmt.Println(*port, *mode)

}
