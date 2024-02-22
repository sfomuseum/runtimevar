package runtimevar

import (
	"context"
	"fmt"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
	"path/filepath"
	"testing"
)

func TestConstantVar(t *testing.T) {

	ctx := context.Background()

	tests := map[string]string{
		"constant://?val=hello+world": "hello world",
		"/foo/bar/baz":                "/foo/bar/baz",
		"test":                        "test",
	}

	for uri, expected := range tests {

		str_var, err := StringVar(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to retrieve string var for '%s', %v", uri, err)
		}

		if str_var != expected {
			t.Fatalf("Unexpected result for '%s'. Expected '%s' but got '%s'", uri, expected, str_var)
		}
	}

}

func TestFileVar(t *testing.T) {

	ctx := context.Background()

	tests := map[string]string{
		"fixtures/helloworld.txt": "hello world\n",
	}

	for rel_path, expected := range tests {

		abs_path, err := filepath.Abs(rel_path)

		if err != nil {
			t.Fatalf("Failed to determine absolute path for '%s', %v", rel_path, err)
		}

		uri := fmt.Sprintf("file://%s?decoder=string", abs_path)

		str_var, err := StringVar(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to retrieve string var for '%s', %v", uri, err)
		}

		if str_var != expected {
			t.Fatalf("Unexpected result for '%s'. Expected '%s' but got '%s'", uri, expected, str_var)
		}
	}

}
