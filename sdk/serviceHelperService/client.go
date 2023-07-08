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

func (ServiceHelperClient *ServiceHelperClient) GetServices(StoId string, CustomSelect map[string]interface{}) ([]Service, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "get",
		Data: map[string]interface{}{
			"sto_id": StoId,
			"select": CustomSelect,
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
