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

func TestRESTapiCreateUser(t *testing.T) {
	var user = &pb.User{
		Username: "nguyendt456",
		Password: "123456789",
		Usertype: "backofficer",
	}
	response := &pb.Response{}

	res := SendHttpRequest(Create_user_gw_port+"/v1/registry", user, &map[string]string{"token": "123"})
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	log.Print(response)
}
