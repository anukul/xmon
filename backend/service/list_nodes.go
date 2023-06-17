package service

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/anukul/xmon/backend/proto"
)

func (s *Service) ListNodes(context.Context, *connect_go.Request[proto.ListNodesRequest]) (*connect_go.Response[proto.ListNodesResponse], error) {
	var nodes []*proto.ListNodesResponse_Node
	for _, node := range s.monitor.Nodes {
		nodes = append(nodes, &proto.ListNodesResponse_Node{
			Name: node.Name,
			ExecutionClient: &proto.ListNodesResponse_Client{
				Version:  node.Execution.Version,
				Status:   node.Execution.Status,
				LastPing: timestamppb.New(node.Execution.LastPing),
			},
		})
	}
	return connect_go.NewResponse(&proto.ListNodesResponse{Nodes: nodes}), nil
}
