package postService

import (
	"encoding/json"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type PostClient struct {
	Client sdk.Client
}

func Client() *PostClient {
	return &PostClient{
		Client: sdk.Client{Url: "http://host.docker.internal:8680"},
	}
}

func (PostClient *PostClient) FindPost(ServiceId string) (*Post, error) {
	Select := map[string]interface{}{
		"internal_id": ServiceId,
	}

	Services, Err := PostClient.GetPosts(Select)
	if Err != nil {
		return nil, Err
	}

	// @TODO Add validation if sto not found
	return &Services[0], nil
}

func (PostClient *PostClient) GetPosts(CustomSelect map[string]interface{}) ([]Post, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "index",
		Data: map[string]interface{}{
			"select": CustomSelect,
		},
	}

	ClientResponseBody, Err := PostClient.Client.SendPostRequest(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	var Posts []Post

	ServicesJson, Err := json.Marshal(ClientResponseBody.Result)
	if Err != nil {
		return nil, Err
	}

	Err = json.Unmarshal(ServicesJson, &Posts)
	if Err != nil {
		return nil, Err
	}

	return Posts, nil
}
