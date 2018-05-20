package main

import (
	rolesPb "github.com/sluongng/grpc-example/role-service/pb"
	"github.com/sluongng/grpc-example/user-service/pb"
	"google.golang.org/grpc"
	"log"
	"context"
	"errors"
	"net"
)

type Server struct {
	users []*pb.User
	rolesClient rolesPb.RolesClient
}

func (s *Server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.UserReply, error) {
	if req.UserId < 0 || req.UserId > int32(len(s.users)) {
		return nil, errors.New("invalid user")
	}

	user := s.users[req.UserId]
	roleReq := &rolesPb.GetUserRoleRequest{
		UserId: req.UserId,
	}
	rolesReply, err := s.rolesClient.GetUserRoles(context.Background(), roleReq)
	if err != nil {
		return nil, err
	}

	roles := make([]*pb.Role, 0)
	for _, r := range rolesReply.Roles {
		roles = append(roles, &pb.Role{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	return &pb.UserReply{
		User: user,
		Roles: roles,
	}, nil
}

func getRolesClient() rolesPb.RolesClient {
	conn, err := grpc.Dial("localhost:7010", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to start gRPC connection: %v", err)
	}

	rolesPb.NewRolesClient(conn)
}

func main() {
	users := []*pb.User{
		{
			Id: 1,
			Email: "alice@email.com",
			Name: "Alice",
		},
		{
			Id: 2,
			Email: "bob@email.com",
			Name: "Bob",
		},
		{
			Id: 3,
			Email: "charly@email.com",
			Name: "Charly",
		},
		{
			Id: 4,
			Email: "dorton@email.com",
			Name: "Dorton",
		},
		{
			Id: 5,
			Email: "edward@email.com",
			Name: "Edward",
		},
	}

	lis, err := net.Listen("tcp", "localhost:7010")
	if err != nil {
		log.Fatal("failed to initialize TCP listener: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	roleServer := &Server{
		users: users,
		rolesClient: getRolesClient(),
	}
	pb.RegisterUsersServer(server, roleServer)

	server.Serve(lis)
}
