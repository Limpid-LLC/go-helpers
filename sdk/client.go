package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	Url      string
	Metadata interface{} `json:"metadata"`
}

type ClientRequestBody struct {
	Method   string      `json:"method"`
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata"`
}

type ClientResponseBody struct {
	Result any `json:"result"`
}

func (Client *Client) SendPostRequest(ClientRequestBody ClientRequestBody) (*ClientResponseBody, error) {
	ClientRequestBody.Metadata = Client.Metadata

	// Encode request struct to json string
	ClientRequestBodyJson, Err := json.Marshal(ClientRequestBody)
	if Err != nil {
		return nil, Err
	}

	// Make request to the service
	Request, Err := http.NewRequest("POST", Client.Url, bytes.NewBuffer(ClientRequestBodyJson))
	if Err != nil {
		return nil, Err
	}

	// Do request and select response
	Response, Err := http.DefaultClient.Do(Request)
	if Err != nil {
		return nil, Err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(Response.Body)

	// Make response struct
	var ClientResponseBody ClientResponseBody

	Err = json.NewDecoder(Response.Body).Decode(&ClientResponseBody)
	if Err != nil {
		return nil, Err
	}

	return &ClientResponseBody, nil
}
