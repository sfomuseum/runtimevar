package runtimevar

import (
	"context"
	gc "gocloud.dev/runtimevar"
	"net/url"
)

func StringVar(ctx context.Context, uri string) (string, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return "", err
	}

	q := u.Query()

	if q.Get("decoder") == "" {
		q.Set("decoder", "string")
		u.RawQuery = q.Encode()
	}

	uri = u.String()

	v, err := gc.OpenVariable(ctx, uri)

	if err != nil {
		return "", err
	}

	defer v.Close()

	snapshot, err := v.Latest(ctx)

	if err != nil {
		return "", err
	}

	return snapshot.Value.(string), nil
}
