package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

type UserTestSuite struct {
	suite.Suite
}

func (suite *UserTestSuite) Test1_TestRESTapiCreateUser() {
	var user = &pb.User{
		Username: "nguyendt456",
		Password: "123456789",
		Usertype: "backofficer",
	}
	response := &pb.Response{}

	res := SendHttpRequest("http://"+setup_env.Create_user_gw+"/v1/registry", user, nil)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	log.Println(response)
	assert.Equal(suite.T(), 202, response.StatusCode)
}
func (suite *UserTestSuite) Test2_RESTapiLoginUser() {
	var user = &pb.LoginForm{
		Username: "nguyendt456",
		Password: "123456789",
	}
	response := &pb.LoginResponse{}

	res := SendHttpRequest("http://"+setup_env.Login_user_gw+"/v1/login", user, nil)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)

}

func TestMain(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
