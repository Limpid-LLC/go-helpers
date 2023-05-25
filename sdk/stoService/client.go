package stoService

import (
	"encoding/json"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type StoClient struct {
	Client sdk.Client
}

func Client() *StoClient {
	return &StoClient{
		Client: sdk.Client{Url: "http://devops.f1xiq.com:9180"},
	}
}

func (StoClient *StoClient) Find(StoId string) (*Sto, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "index",
		Data: map[string]interface{}{
			"select": map[string]interface{}{
				"internal_id": StoId,
			},
		},
	}

	Stos, Err := StoClient.Get(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	// @TODO Add validation if sto not found
	return &Stos[0], nil
}

func (StoClient *StoClient) Get(ClientRequestBody sdk.ClientRequestBody) ([]Sto, error) {
	ClientResponseBody, Err := StoClient.Client.SendPostRequest(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	var Stos []Sto

	StosJson, Err := json.Marshal(ClientResponseBody.Result)
	if Err != nil {
		return nil, Err
	}

	Err = json.Unmarshal(StosJson, &Stos)
	if Err != nil {
		return nil, Err
	}

	return Stos, nil
}
