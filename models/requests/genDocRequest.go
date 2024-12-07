package requests

type GenDocRequest struct {
	MrLinks     []string `json:"mrLinks" binding:"required,min=1"`
	GitlabToken string   `json:"gitlabToken" binding:"required"`
	Model       string   `json:"model" binding:"required"`
}
