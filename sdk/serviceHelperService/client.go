package serviceHelperService

import (
	"encoding/json"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type ServiceHelperClient struct {
	Client sdk.Client
}

func Client(metadata interface{}) *ServiceHelperClient {
	return &ServiceHelperClient{
		Client: sdk.Client{
			Url:      "http://devops.f1xiq.com:8880",
			Metadata: metadata,
		},
	}
}

func New(url string) *ServiceHelperClient {
	return &ServiceHelperClient{
		Client: sdk.Client{
			Url: url,
		},
	}
}

func (ServiceHelperClient *ServiceHelperClient) SetMetadata(metadata interface{}) {
	ServiceHelperClient.Client.Metadata = metadata
}

func (ServiceHelperClient *ServiceHelperClient) GetServices(CustomSelect map[string]interface{}, stoId string) ([]Service, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "get",
		Data: map[string]interface{}{
			"select":        CustomSelect,
			"sto_id":        stoId,
			"without_group": true,
		},
	}

	ClientResponseBody, Err := ServiceHelperClient.Client.SendPostRequest(ClientRequestBody)
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
