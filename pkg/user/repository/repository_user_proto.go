package repository

import (
	"context"
	"go_services_lab/pkg/user/proto"
)

type GRPCUsersServer struct {
	User
}

func (s *GRPCUsersServer) IsExistById(ctx context.Context, req *proto.IsExistByIdRequest) (*proto.IsExistByIdResponse, error) {
	_, err := s.Get(int(req.Id))

	if err != nil {
		return &proto.IsExistByIdResponse{IsExist: false}, nil
	}

	return &proto.IsExistByIdResponse{IsExist: true}, nil
}
