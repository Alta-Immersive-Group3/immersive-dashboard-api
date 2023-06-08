package handler

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"
)

type AuthResponse struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

type UserResponse struct {
	Id       uint64 `json:"id,omitempty"`
	FullName string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	IdTeam   uint64 `json:"id_team,omitempty"`
	Status   bool   `json:"status,omitempty"`
}

func CoreToAuthResponse(user user.Core, jwtToken string) AuthResponse {
	return AuthResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: jwtToken,
		Role:  user.Role,
	}
}

func CoreToGetUserResponse(user user.Core) UserResponse {
	return UserResponse{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		IdTeam:   user.IdTeam,
		Status:   user.Status,
	}
}
