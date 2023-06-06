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
	Id        uint   `json:"id" form:"id"`
	FullName  string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Role      string `json:"role" form:"role"`
	IdTeam    uint64 `json:"id_team" form:"id_team"`
	Status    bool   `json:"status" form:"status"`
	IsDeleted bool   `json:"is_deleted" form:"is_deleted"`
}

func CoreToAuthResponse(user user.Core, jwtToken string) AuthResponse {
	return AuthResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: jwtToken,
		Role:  user.Role,
	}
}
