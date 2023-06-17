package service

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/anukul/xmon/backend/proto"
)

func (s *Service) ListSlots(context.Context, *connect_go.Request[proto.ListSlotsRequest]) (*connect_go.Response[proto.ListSlotsResponse], error) {
	return nil, nil
}
