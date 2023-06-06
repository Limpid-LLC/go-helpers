package postService

import (
	"encoding/json"
	"errors"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type PostClient struct {
	Client sdk.Client
}

func Client() *PostClient {
	return &PostClient{
		Client: sdk.Client{Url: "http://devops.f1xiq.com:8680"},
	}
}

func (PostClient *PostClient) FindPost(ServiceId string) (*Post, error) {
	Select := map[string]interface{}{
		"internal_id": ServiceId,
	}

	Posts, Err := PostClient.GetPosts(Select)
	if Err != nil {
		return nil, Err
	}

	if len(Posts) == 0 {
		return nil, errors.New("post not found")
	}

	return &Posts[0], nil
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
