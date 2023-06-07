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
	Id       uint64 `json:"id" form:"id"`
	FullName string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
	IdTeam   uint64 `json:"id_team" form:"id_team"`
	Status   bool   `json:"status" form:"status"`
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
