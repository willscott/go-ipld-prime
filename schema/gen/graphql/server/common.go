package gengraphqlserver

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
	"github.com/warpfork/go-wish"
)

func EmitFileHeader(w io.Writer, pkg, tsPkg string, c *config) {
	writeTemplate(`
	// Code generated by github.com/ipld/go-ipld-prime/schema/gen/graphql/server, DO NOT EDIT.

	package {{ .Package }}

	import (
		"context"
		"fmt"

		"github.com/graphql-go/graphql"
		"github.com/graphql-go/graphql/language/ast"
		ipld "github.com/ipld/go-ipld-prime"
		cidlink "github.com/ipld/go-ipld-prime/linking/cid"
		"{{ .TSPackage }}"
	)

	type nodeLoader func(ctx context.Context, cid cidlink.Link, builder ipld.NodeBuilder) error
	const nodeLoaderCtxKey = "NodeLoader"

	`, w, map[string]string{
		"Package":   pkg,
		"TSPackage": tsPkg,
	}, c)
}

func EmitFileCompletion(w io.Writer, ts schema.TypeSystem, c *config) {
	w.Write([]byte("\nfunc init() {\n"))
	w.Write([]byte(wish.Dedent(string(c.initDirectives.Bytes()))))
	w.Write([]byte("\n}\n"))
}