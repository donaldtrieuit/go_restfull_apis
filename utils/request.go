package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/

var client = &http.Client{}

type HttpResponse struct {
	StatusCode int
	Body       []byte
}

func POST(url string, body []byte, header http.Header) (*HttpResponse, error) {
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("- Request url: POST %s \n- Body : %s \n", url, string(body))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		readBody, err := ioutil.ReadAll(resp.Body)

		defer resp.Body.Close()

		if err != nil {
			return nil, err
		}

		fmt.Printf("- Resposne: %s \n", string(readBody))
		return &HttpResponse{StatusCode: resp.StatusCode, Body: readBody}, err
	}
	return nil, err
}

func GET(url string, header http.Header) (*HttpResponse, error) {
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("- Request url: GET %s \n- Body : %s \n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		readBody, err := ioutil.ReadAll(resp.Body)

		defer resp.Body.Close()

		if err != nil {
			return nil, err
		}

		fmt.Printf("- Resposne: %s \n", string(readBody))
		return &HttpResponse{StatusCode: resp.StatusCode, Body: readBody}, err
	}
	return nil, err
}

func PUT(url string, body []byte, header http.Header) (*HttpResponse, error) {
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("- Request url: PUT %s \n- Body : %s \n", url, string(body))

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		readBody, err := ioutil.ReadAll(resp.Body)

		defer resp.Body.Close()

		if err != nil {
			return nil, err
		}

		fmt.Printf("- Resposne: %s \n", string(readBody))
		return &HttpResponse{StatusCode: resp.StatusCode, Body: readBody}, err
	}
	return nil, err
}

func DELETE(url string, body []byte, header http.Header) (*HttpResponse, error) {
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("- Request url: DELETE %s \n- Body : %s \n", url, string(body))

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		readBody, err := ioutil.ReadAll(resp.Body)

		defer resp.Body.Close()

		if err != nil {
			return nil, err
		}

		fmt.Printf("- Resposne: %s \n", string(readBody))
		return &HttpResponse{StatusCode: resp.StatusCode, Body: readBody}, err
	}
	return nil, err
}

func GetHttpResponse(data interface{}) HttpResponse {
	res := HttpResponse{}
	bodyBytes, _ := json.Marshal(data)
	_ = json.Unmarshal(bodyBytes, &res)
	return res
}

func GetArrayHttpResponse(data interface{})[]HttpResponse {
	var res []HttpResponse
	bodyBytes, _ := json.Marshal(data)
	_ = json.Unmarshal(bodyBytes, &res)
	return res
}