package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/font/proto/font"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"unsafe"
)

var (
	service = grpc.NewService(
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)
)

func TestClient(t *testing.T) {
	srv := font.NewFontsService("go.micro.srv.font", service.Client())
	rsp, err := srv.Index(context.Background(), &font.IndexRequest{
		Page:     0,
		PageSize: 1,
		Status:   1,
	})

	if err != nil {
		t.Error(err)
	}


	if len(rsp.Data) != 1 {
		t.Error("rep Data len more than 1 ")
	}
}

func BenchmarkGRPCClient_FontService(b *testing.B) {
	srv := font.NewFontsService("go.micro.srv.font", service.Client())
	for i := 0; i < b.N; i++ {
		rsp, err := srv.Index(context.Background(), &font.IndexRequest{
			Page:     0,
			PageSize: 10000,
			Status:   1,
		})
		if err != nil {
			b.Error(err)
		}

		if len(rsp.Data) < 1 {
			b.Error("rep Data len more than 1 ")
		}
	}
}

func TestApiClient(t *testing.T) {
	json := `{"service":"go.micro.srv.font","method":"Fonts.Index","request":{"page":0,"pageSize":1000,"status":1}}`
	reader := bytes.NewReader([]byte(json))
	url := "http://localhost:8082/rpc"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	if len(*str) < 1000 {
		t.Error("长度不对")
	}
}

func BenchmarkApiClient_FontService(b *testing.B) {
	json := `{"service":"go.micro.srv.font","method":"Fonts.Index","request":{"page":0,"pageSize":1000,"status":1}}`
	reader := bytes.NewReader([]byte(json))
	url := "http://localhost:8082/rpc"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}

	for i := 0; i < b.N; i++ {
		resp, err := client.Do(request)
		if err != nil {
			b.Error(err.Error())
		}
		//respBytes, err := ioutil.ReadAll(resp.Body)
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Error(err.Error())
		}

		//byte数组直接转成string，优化内存
		//str := (*string)(unsafe.Pointer(&respBytes))
		//if len(*str) < 1000 {
		//	b.Error("长度不对")
		//}
	}

}
