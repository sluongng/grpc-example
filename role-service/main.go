package main

import (
	pb "github.com/sluongng/grpc-example/role-service/pb"
	"net"
	"log"
	"google.golang.org/grpc"
	"context"
	"errors"
)

type Server struct {
	userRoles map[int32][]*pb.Role
	roles     []*pb.Role
}

func (s *Server) GetRoles(_ context.Context, _ *pb.EmptyRequest) (*pb.RolesReply, error) {
	return &pb.RolesReply{
		Roles: s.roles,
	}, nil
}

func (s *Server) GetUserRole(_ context.Context, req *pb.GetUserRoleRequest) (*pb.UserRoleReply, error) {
	roles, ok := s.userRoles[req.UserId]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &pb.UserRoleReply{
		UserId: req.UserId,
		Roles:  roles,
	}, nil
}

func main() {
	var (
		normal = &pb.Role{
			Id:   1,
			Name: "normal",
		}
		manager = &pb.Role{
			Id:   1,
			Name: "manager",
		}
		cto = &pb.Role{
			Id:   1,
			Name: "CTO",
		}
		ceo = &pb.Role{
			Id:   1,
			Name: "CEO",
		}
	)

	lis, err := net.Listen("tcp", "localhost:7060")
	if err != nil {
		log.Fatalf("failed to initialize TCP listerner: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	roleServer := &Server{
		userRoles: map[int32][]*pb.Role{
			1: {normal},
			2: {normal, manager},
			3: {manager},
			4: {cto},
			5: {ceo},
		},
		roles: []*pb.Role{normal, manager, cto, ceo},
	}
	pb.RegisterRolesServer(server, roleServer)

	server.Serve(lis)
}
