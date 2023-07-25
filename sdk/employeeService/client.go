package employeeService

import (
	"encoding/json"
	"github.com/Limpid-LLC/go-helpers/sdk"
)

type EmployeeClient struct {
	Client sdk.Client
}

func Client(metadata interface{}) *EmployeeClient {
	return &EmployeeClient{
		Client: sdk.Client{
			Url:      "http://devops.f1xiq.com:8980",
			Metadata: metadata,
		},
	}
}

func (PostClient *EmployeeClient) GetEmployees(CustomSelect map[string]interface{}) ([]Employee, error) {
	ClientRequestBody := sdk.ClientRequestBody{
		Method: "get",
		Data: map[string]interface{}{
			"select": CustomSelect,
		},
	}

	ClientResponseBody, Err := PostClient.Client.SendPostRequest(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	var Employees []Employee

	ServicesJson, Err := json.Marshal(ClientResponseBody.Result)
	if Err != nil {
		return nil, Err
	}

	Err = json.Unmarshal(ServicesJson, &Employees)
	if Err != nil {
		return nil, Err
	}

	return Employees, nil
}
