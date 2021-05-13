package playground

import (
	"context"
	"net"

	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

func SUH() {
	m := drpcmux.New()

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("ha")
	}

	ctx := context.Background()
	drpcserver.New(m).Serve(ctx, lis)
}
