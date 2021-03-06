package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/runtimevar"
	_ "gocloud.dev/runtimevar/awsparamstore"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
	"log"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	for _, uri := range flag.Args() {

		str_var, err := runtimevar.StringVar(ctx, uri)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(str_var)
	}
}
