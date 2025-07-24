package requests

type GenDocRequest struct {
	MrLinks []string `json:"mrLinks" binding:"required,min=1"`
	Model   string   `json:"model" binding:"required"`
}
