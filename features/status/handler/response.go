package handler

type StatusResponse struct {
	Id   uint64 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
