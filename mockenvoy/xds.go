package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/wzshiming/xds/utils"
	"google.golang.org/grpc"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var (
	address = "istiod.istio-system.svc:15010"
	nodeId  = ""
)

func init() {
	flag.StringVar(&address, "u", address, "xds server")
	flag.StringVar(&nodeId, "n", nodeId, "node id")
	flag.Parse()
}

var jsonpbMarshaler = jsonpb.Marshaler{
	AnyResolver: dynamicAnyResolver{},
}

type dynamicAnyResolver struct {
}

func (dynamicAnyResolver) Resolve(typeURL string) (proto.Message, error) {
	mt, err := protoregistry.GlobalTypes.FindMessageByURL(typeURL)
	if err != nil {
		log.Println(err, typeURL)
		return &empty.Empty{}, nil
	}
	return proto.MessageV1(mt.New().Interface()), nil
}

func show(m proto.Message) {
	data, err := jsonpbMarshaler.MarshalToString(m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

func xdsServer() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// todo: use agg client

	xds := discovery.NewAggregatedDiscoveryServiceClient(conn)
	stm, err := xds.StreamAggregatedResources(context.Background())
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	ctx := stm.Context()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := ctx.Err()
			if err != nil {
				panic(err)
			}

			msg, err := stm.Recv()
			if err != nil {
				panic(err)
			}

			for _, rsc := range msg.Resources {
				// fmt.Println("rsc", rsc.TypeUrl)
				switch rsc.TypeUrl {
				case resource.ListenerType:
					ll := &envoy_config_listener_v3.Listener{}
					_ = proto.Unmarshal(rsc.Value, ll)
					show(ll)
				}
			}
		}
	}()

	node := &utils.NodeConfig{NodeID: nodeId}

	err = stm.Send(&discovery.DiscoveryRequest{
		Node: &core.Node{
			Id:       node.ID(),
			Metadata: node.Meta(),
		},
		TypeUrl:       resource.ListenerType,
		VersionInfo:   "",
		ResourceNames: nil,
	})
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
