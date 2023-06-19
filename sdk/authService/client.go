package authService

import (
	"encoding/json"
	"errors"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type AuthClient struct {
	Client sdk.Client
}

func Client() *AuthClient {
	return &AuthClient{
		Client: sdk.Client{Url: "http://devops.f1xiq.com:9001"},
	}
}

func (AuthClient *AuthClient) FindRole(data map[string]any) (*Role, error) {
	Roles, Err := AuthClient.GetRoles(data)
	if Err != nil {
		return nil, Err
	}

	if len(Roles) == 0 {
		return nil, errors.New("role not found")
	}

	return &Roles[0], nil
}

func (AuthClient *AuthClient) GetRoles(data map[string]any) ([]Role, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "get_roles",
		Data:   data,
	}

	ClientResponseBody, Err := AuthClient.Client.SendPostRequest(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	var Roles []Role

	StosJson, Err := json.Marshal(ClientResponseBody.Result)
	if Err != nil {
		return nil, Err
	}

	Err = json.Unmarshal(StosJson, &Roles)
	if Err != nil {
		return nil, Err
	}

	return Roles, nil
}
