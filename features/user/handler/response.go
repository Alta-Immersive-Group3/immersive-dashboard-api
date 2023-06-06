package handler

import "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"

type AuthResponse struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

func CoreToAuthResponse(user user.Core, jwtToken string) AuthResponse {
	return AuthResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: jwtToken,
		Role:  user.Role,
	}
}
