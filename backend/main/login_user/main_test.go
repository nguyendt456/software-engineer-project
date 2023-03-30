package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
)

func SendHttpRequest(url string, payload interface{}, headers *map[string]string) *http.Response {
	request_body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Cannot convert to JSON:", err)
		return nil
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(request_body))
	if err != nil {
		log.Fatal("Error when request:", err)
		return nil
	}
	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}
	res, _ := http.DefaultClient.Do(req)

	return res
}

func TestRESTapiLoginUser(t *testing.T) {
	var user = &pb.LoginForm{
		Username: "nguyendt456",
		Password: "123456789",
	}
	response := &pb.LoginResponse{}

	res := SendHttpRequest("http://"+setup_env.Login_user_gw+"/v1/login", user, nil)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Println(response)
}
