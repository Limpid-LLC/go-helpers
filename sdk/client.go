package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	Url   string
	Token string
}

type ClientRequestBody struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

type ClientResponseBody struct {
	Result []map[string]interface{} `json:"result"`
}

func (Client *Client) SendPostRequest(ClientRequestBody ClientRequestBody) (*ClientResponseBody, error) {
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

	// Add auth token to request
	Request.Header.Set("Token", Client.Token)

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
