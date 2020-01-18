package graph

import (
	"context"
	"fmt"

	"../../cmd"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var (
	connHost string
	connPort int
	conn     *grpc.ClientConn
	dg       *dgo.Dgraph
	ctx      context.Context
)

// Connect connects to the database
func Connect() bool {
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", connHost, connPort), grpc.WithInsecure())
	if err != nil {
		l.Errorf("Unable to connect to DGraph: %s", err)
		return false
	}
	dc := api.NewDgraphClient(conn)
	dg = dgo.NewDgraphClient(dc)
	ctx = context.Background()
	return true
}

func init() {
	cmd.GetString("db.dgraph.host", &connHost)
	cmd.GetInt("db.dgraph.port", &connPort)
}
