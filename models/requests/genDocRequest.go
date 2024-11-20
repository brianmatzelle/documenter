package requests

type GenDocRequest struct {
	MrLink      string `json:"mrLink"`
	GitlabToken string `json:"gitlabToken"`
	Model       string `json:"model"`
}
