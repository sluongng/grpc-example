package user_service

import (
	"v/github.com/gorilla/mux@v1.6.2"
	"awesomeProject/user-service/pb"
)

type Server struct {
	roles []*pb.Role
}

func main() {
	r := mux.NewRouter()
	_ = r
}
