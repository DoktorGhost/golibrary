package grpcProvider

import (
	"context"
	"github.com/DoktorGhost/golibrary/users/internal/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/users/internal/services"
	pb "github.com/DoktorGhost/golibrary/users/pkg/proto"
)

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	repo *services.UserService
}

func NewUserGRPCServer(us *services.UserService) *UsersServiceServer {
	return &UsersServiceServer{repo: us}
}

func (s *UsersServiceServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetUserById(int(req.Id))
	if err != nil {
		return &pb.GetUserResponse{Error: err.Error()}, err
	}
	return &pb.GetUserResponse{
		Id:           int32(user.ID),
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		FullName:     user.FullName,
	}, nil
}

func (s *UsersServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	err := s.repo.UpdateUser(dao.UserTable{
		ID:           int(req.Id),
		Username:     req.Username,
		PasswordHash: req.PasswordHash,
		FullName:     req.FullName,
	})
	if err != nil {
		return &pb.UpdateUserResponse{Error: err.Error()}, err
	}
	return &pb.UpdateUserResponse{}, nil
}
func (s *UsersServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.repo.DeleteUser(int(req.Id))
	if err != nil {
		return &pb.DeleteUserResponse{Error: err.Error()}, nil
	}
	return &pb.DeleteUserResponse{}, nil
}
func (s *UsersServiceServer) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil {
		return &pb.GetUserResponse{Error: err.Error()}, err
	}
	return &pb.GetUserResponse{
		Id:           int32(user.ID),
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		FullName:     user.FullName,
	}, nil
}
