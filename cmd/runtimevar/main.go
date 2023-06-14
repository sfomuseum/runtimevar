package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sfomuseum/runtimevar"
	_ "gocloud.dev/runtimevar/awsparamstore"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
)

func main() {

	timeout := flag.Int("timeout", 0, "The maximum number of second in which a variable can be resolved. If 0 no timeout is applied.")

	flag.Parse()

	ctx := context.Background()

	if *timeout > 0 {

		c, cancel := context.WithTimeout(ctx, time.Duration(*timeout)*time.Second)
		defer cancel()

		ctx = c
	}

	for _, uri := range flag.Args() {

		str_var, err := runtimevar.StringVar(ctx, uri)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(str_var)
	}
}
