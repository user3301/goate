package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/user3301/grpclab/pkg/proto"
	"github.com/user3301/grpclab/pkg/store"
	"github.com/user3301/grpclab/pkg/types"
)

type UserService struct {
	userStore store.UserStorer
}

func NewUserService(userStore store.UserStorer) *UserService {
	return &UserService{userStore: userStore}
}

func (u UserService) Register(ctx context.Context, credentials *proto.Credentials) (*proto.RegisterResponse, error) {
	userDetails := types.UserDetails{
		Username: credentials.GetUsername(),
		Password: credentials.GetPassword(),
	}
	err := u.userStore.CreateUser(ctx, userDetails)
	return &proto.RegisterResponse{}, err
}

func (u UserService) Login(ctx context.Context, _ *proto.Credentials) (*proto.LoginResponse, error) {
	log.Print("user service entered")
	var username, password string
	var err error
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		authHeader := md.Get("authorization")
		if len(authHeader) != 1 || len(authHeader[0]) < 7 {
			return nil, fmt.Errorf("no basic auth")
		}
		basicAuth := authHeader[0][6:]
		username, password, err = parseBasicAuth(basicAuth)
		if err != nil {
			return nil, err
		}
	}
	userDetails := types.UserDetails{
		Username: username,
		Password: password,
	}
	verified, reason := u.userStore.Verify(ctx, userDetails)
	if verified {
		return &proto.LoginResponse{}, nil
	}
	return nil, fmt.Errorf(reason)
}

func (u UserService) ChangePassword(ctx context.Context,
	request *proto.ChangePasswordRequest) (*proto.ChangePasswordResponse, error) {
	log.Printf("change password entered %v", request)
	userDetails := types.UserDetails{
		Username: request.GetUsername(),
		Password: request.GetOldPassword(),
	}
	verified, reason := u.userStore.Verify(ctx, userDetails)
	if !verified {
		return nil, fmt.Errorf(reason)
	}
	newDetails := types.UserDetails{
		Username: request.GetUsername(),
		Password: request.GetNewPassword(),
	}
	updated, err := u.userStore.UpdateUser(ctx, newDetails)
	if err != nil || !updated {
		return nil, err
	}
	return &proto.ChangePasswordResponse{}, nil
}

func parseBasicAuth(auth string) (username, password string, err error) {
	c, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		err = fmt.Errorf("invalid basic auth header")
		return
	}
	return cs[:s], cs[s+1:], nil
}
