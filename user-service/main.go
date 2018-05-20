package user_service

import (
	"v/github.com/gorilla/mux@v1.6.2"
	pb "awesomeProject/user-service/pb"
)

type Server struct {
	users []*pb.User
}

func main() {
	r := mux.NewRouter()
	_ = r
}
