package handler

import "github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/user"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	FullName string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	IdTeam   uint64 `json:"id_team" form:"id_team"`
}

func UserRequestToCore(userRequest UserRequest) user.Core {
	return user.Core{
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     userRequest.Role,
		IdTeam:   userRequest.IdTeam,
	}
}
