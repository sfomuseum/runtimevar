# runtimevar

Simple wrapper around the Go Cloud runtimevar package

## Example

```
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/runtimevar"
	_ "gocloud.dev/runtimevar/awsparamstore"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	for _, uri := range flag.Args() {
		str_var, _ := runtimevar.StringVar(ctx, uri)
		fmt.Printf(str_var)
	}
}
```

## Tools

### runtimevar

```
go run cmd/runtimevar/main.go 'constant://?val=hello+world'
hello world
```

The following Go Cloud `runtimevar` services are supported by the runtimevar tool:

* [AWS Parameter Store](https://gocloud.dev/howto/runtimevar/#awsps)
* [Blob](https://gocloud.dev/howto/runtimevar/#blob)
* [Local](https://gocloud.dev/howto/runtimevar/#local)

## See also

* https://gocloud.dev/howto/runtimevar/