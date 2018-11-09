package  main

import (
	"os"
	"github.com/argoproj/argo-events/common"
	"fmt"
	"net"
	"google.golang.org/grpc"
	gp "github.com/argoproj/argo-events/gateways/grpc/proto"
	"github.com/argoproj/argo-events/gateways"
	"github.com/argoproj/argo-events/gateways/grpc/calendar"
)

func main() {
	rpcServerPort, ok := os.LookupEnv(common.GatewayProcessorGRPCServerPort)
	if !ok {
		panic("gateway rpc server port is not provided")
	}

	gc := gateways.NewGatewayConfiguration()
	ce := &calendar.CalendarConfigExecutor{}
	ce.GatewayConfig = gc

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", rpcServerPort))
	if err != nil {
		ce.Log.Fatal().Err(err).Msg("server failed to listen")
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	gp.RegisterGatewayExecutorServer(grpcServer, ce)
	ce.Log.Info().Str("port", rpcServerPort).Msg("gRPC server started listening...")
	grpcServer.Serve(lis)
}

