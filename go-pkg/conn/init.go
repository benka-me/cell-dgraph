package conn

import (
	"context"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

type Dgraph struct {
	*dgo.Dgraph
}

func InitDgraph(dg *dgo.Dgraph, dropAll bool) error {
	op := &api.Operation{}
	op.DropAll = dropAll
	op.Schema = `
		Username: string @index(exact) .
		Email: string @index(exact) .
		Password: password .
		Firstname: string @index(term, fulltext) .
		Lastname: string @index(term, fulltext) .
		Age: int @index(int) .
		Owner: bool @index(bool) .

		CustomerName: string @index(exact) .
		WelcomeMessage: string .

		Room: string @index(exact) .
	`
	ctx := context.Background()
	return dg.Alter(ctx, op)
}

func NewClient() (*dgo.Dgraph, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	d, err := grpc.Dial("localhost:9080", dialOpts...)

	if err != nil {
		return nil, nil
	}

	ret := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	err = InitDgraph(ret, false)
	return ret, err
}
