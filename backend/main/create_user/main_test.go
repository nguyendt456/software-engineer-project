package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
)

// func TestGrpcCreateUser(t *testing.T) {
// 	grpc_conn, err := grpc.Dial("0.0.0.0"+":8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		t.Fatal(err.Error())
// 		return
// 	}
// 	defer grpc_conn.Close()

// 	client := pb.NewCreateUserServiceClient(grpc_conn)

// 	response, err := client.CreateUser(context.Background(), &pb.User{
// 		Username: "nguyendt456",
// 		Password: "123456789",
// 		Usertype: "backofficer",
// 	})

// 	if err != nil {
// 		t.Fatal(err.Error())
// 		return
// 	}
// 	assert.Equal(t, 201, int(response.StatusCode))
// 	assert.Equal(t, "User Created", response.Message)
// }

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

func TestRESTapiCreateUser(t *testing.T) {
	var user = &pb.User{
		Username: "nguyendt456",
		Password: "123456789",
		Usertype: "backofficer",
	}
	response := &pb.Response{}

	res := SendHttpRequest("http://"+setup_env.Create_user_gw+"/v1/registry", user, nil)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	fmt.Print(response)
}
