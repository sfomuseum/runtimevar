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
$> go run cmd/runtimevar/main.go 'constant://?val=hello+world'
hello world
```

The following Go Cloud `runtimevar` services are supported by the runtimevar tool:

* [AWS Parameter Store](https://gocloud.dev/howto/runtimevar/#awsps)
* [Blob](https://gocloud.dev/howto/runtimevar/#blob)
* [Local](https://gocloud.dev/howto/runtimevar/#local)

### AWS Parameter Store

It is possible to load runtime variables from AWS Parameter Store using [aaronland/go-aws-session](https://github.com/aaronland/go-aws-session) credential strings. For example:

```
$> go run cmd/runtimevar/main.go 'awsparamstore://hello-world?region=us-west-2&credentials=session'
hello world
```

Valid `aaronland/go-aws-session` credential strings are:

Credentials for AWS sessions are defined as string labels. They are:

| Label | Description |
| --- | --- |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `{AWS_PROFILE_NAME}` | This this profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | This this profile from a user-defined AWS credentials location. |

## See also

* https://gocloud.dev/howto/runtimevar
* https://github.com/aaronland/go-aws-session