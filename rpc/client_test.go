package rpc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/coschain/contentos-go/common/prototype"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/p2p"
	"github.com/coschain/contentos-go/rpc/pb"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var asc grpcpb.ApiServiceClient

func TestMain(m *testing.M) {
	gs, _ := NewGRPCServer(&node.ServiceContext{})
	err := gs.Start(&p2p.Server{})
	if err != nil {
		fmt.Print(err)
	}
	err = gs.RunGateway()
	if err != nil {
		fmt.Print(err)
	}
	defer gs.Stop()

	addr := fmt.Sprintf("127.0.0.1:%d", uint32(8888))
	conn, err := Dial(addr)
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()

	asc = grpcpb.NewApiServiceClient(conn)

	exitCode := m.Run()
	asc = nil
	os.Exit(exitCode)
}

func TestGRPCApi_GetAccountByName(t *testing.T) {
	req := &grpcpb.GetAccountByNameRequest{AccountName: &prototype.AccountName{Value: "Jack"}}
	resp := &grpcpb.AccountResponse{}
	resp, err := asc.GetAccountByName(context.Background(), req)
	if err != nil {
		t.Errorf("GetAccountByName failed: %x", err)
	} else {
		t.Logf("GetAccountByName detail: %s", resp.AccountName)
	}
}

func TestGPRCApi_GetFollowerListByName(t *testing.T) {
	req := &grpcpb.GetFollowerListByNameRequest{}
	resp := &grpcpb.GetFollowerListByNameResponse{}
	resp, err := asc.GetFollowerListByName(context.Background(), req)
	if err != nil {
		t.Errorf("GetFollowerListByName failed: %x", err)
	} else {
		t.Logf("GetFollowerListByName detail: %s", resp.FollowerList)
	}
}

func TestHTTPApi_GetAccountByName(t *testing.T) {
	postValue := "{\"account_name\": {\"value\":\"jack's test info\"}}"
	http_client("POST", "http://127.0.0.1:8080/v1/user/get_account_by_name", postValue)
}

func TestHTTPApi_GetWitnessList(t *testing.T) {
	http_client("GET", "http://127.0.0.1:8080/v1/user/get_witness_list?page=1&size=5", "")
}

func http_client(rtype, url, reqJson string) error {
	req, err := http.NewRequest(rtype, url, bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}