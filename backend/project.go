package main

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	crm "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

func project_list(token *oauth2.Token) (*[]byte, error) {
	ctx := context.Background()
	crmService, err := crm.NewService(ctx, option.WithTokenSource(googleAuth.TokenSource(ctx, token)))
	projectsResponse, err := crmService.Projects.List().Do()
	if err != nil {
		return nil, err
	}
	projects := projectsResponse.Projects
	// TODO: handle non 200 HTTP responses?
	// TODO: handle empty project list
	var projectNames = make([]string, len(projects))
	for i := 0; i < len(projects); i++ {
		projectNames[i] = projects[i].Name
	}
	responseStruct := AuthorizationResponse{Projects: projectNames}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
