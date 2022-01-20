package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*string, error)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
	//rolePermissions domain.RolePermissions
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*string, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}
func NewLoginService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repo}
}
