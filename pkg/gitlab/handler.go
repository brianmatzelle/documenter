package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMrApiLink(mrLink string) string {
	fmt.Println("Getting merge request API link...")
	// TODO: implement
	// https://gitlab.icg360.net/eng/keystone/-/merge_requests/878 -> https://gitlab.icg360.net/api/v4/projects/347/merge_requests/878/changes
	return "https://gitlab.icg360.net/api/v4/projects/347/merge_requests/878/changes"
}

func GetMrInfo(mrApiLink string, gitlabToken string) (json.RawMessage, error) {
	fmt.Println("Getting merge request info for ", mrApiLink)
	client := &http.Client{}
	req, err := http.NewRequest("GET", mrApiLink, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", gitlabToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
