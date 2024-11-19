package services

import (
	"documenter/models/requests"
	"encoding/json"
	"fmt"

	"documenter/pkg/generate"
	"documenter/pkg/gitlab"
)

func GenerateDocService(request requests.GenDocRequest) (string, error) {
	fmt.Println("Generating document...")
	// get request to request.MrLink, set authentication header with request.GitlabToken
	mrApiLink := gitlab.GetMrApiLink(request.MrLink)
	mrInfo, err := gitlab.GetMrInfo(mrApiLink, request.GitlabToken)
	if err != nil {
		return "", err
	}

	// for verbose
	mrTitle, err := showMrInfoTitle(mrInfo)
	if err != nil {
		return "", err
	}
	fmt.Println("Got merge request info from gitlab, title: ", mrTitle)
	// end for verbose

	// call ollama to generate doc
	fmt.Println("Generating document...")
	docStr, err := generate.GenerateDoc(mrInfo)
	if err != nil {
		return "", err
	}
	fmt.Println("Generated document: ", docStr)
	return docStr, nil
}

func showMrInfoTitle(mrInfo json.RawMessage) (string, error) {
	var mrInfoMap map[string]interface{}
	err := json.Unmarshal(mrInfo, &mrInfoMap)
	if err != nil {
		fmt.Println("Error unmarshalling merge request info: ", err)
		return "", err
	}
	return mrInfoMap["title"].(string), nil
}
