package categoryService

import (
	"encoding/json"
	"errors"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type CategoryClient struct {
	Client sdk.Client
}

func Client(metadata interface{}) *CategoryClient {
	return &CategoryClient{
		Client: sdk.Client{
			Url:      "http://devops.f1xiq.com:8580",
			Metadata: metadata,
		},
	}
}

func New(url string) *CategoryClient {
	return &CategoryClient{
		Client: sdk.Client{
			Url: url,
		},
	}
}

func (CategoryClient *CategoryClient) SetMetadata(metadata interface{}) {
	CategoryClient.Client.Metadata = metadata
}

func (CategoryClient *CategoryClient) FindService(ServiceId string) (*Service, error) {
	Select := map[string]interface{}{
		"type":        "service",
		"internal_id": ServiceId,
	}

	Services, Err := CategoryClient.GetServices(Select)
	if Err != nil {
		return nil, Err
	}

	if len(Services) == 0 {
		return nil, errors.New("post not found")
	}

	return &Services[0], nil
}

func (CategoryClient *CategoryClient) GetServices(CustomSelect map[string]interface{}) ([]Service, error) {
	FunctionSelect := map[string]interface{}{
		"type": "service",
	}

	ClientRequestBody := sdk.ClientRequestBody{
		Method: "index_service",
		Data: map[string]interface{}{
			"select": mergeMaps(FunctionSelect, CustomSelect),
		},
	}

	ClientResponseBody, Err := CategoryClient.Client.SendPostRequest(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	var Services []Service

	ServicesJson, Err := json.Marshal(ClientResponseBody.Result)
	if Err != nil {
		return nil, Err
	}

	Err = json.Unmarshal(ServicesJson, &Services)
	if Err != nil {
		return nil, Err
	}

	return Services, nil
}

func mergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	merged := make(map[K]V)

	for key, value := range m1 {
		merged[key] = value
	}

	for key, value := range m2 {
		merged[key] = value
	}

	return merged
}
