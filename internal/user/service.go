package user

import "context"

type Service struct {
	Storage Storage
	// TODO loger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	return User{}, nil
}
