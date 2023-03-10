package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	pb "github.com/nguyendt456/software-engineer-project/pb"
)

func SendHttpRequest(port string, payload interface{}, headers *map[string]string) *http.Response {
	request_body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Cannot convert to JSON:", err)
		return nil
	}
	req, err := http.NewRequest(http.MethodPost, "http://0.0.0.0"+port, bytes.NewBuffer(request_body))
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
	response := &pb.Response{}

	res := SendHttpRequest(Login_user_gw_port+"/v1/login", user, &map[string]string{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6Im5ndXllbmR0NDU2IiwiVXNlcnR5cGUiOiJiYWNrb2ZmaWNlciIsIk5hbWUiOiIiLCJleHAiOjE2Nzg0NzAzOTZ9._amiWwbRfIocDXRNwziLILYHinkIRO113PWoKt1EH_c"})
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	log.Print(response)
}
