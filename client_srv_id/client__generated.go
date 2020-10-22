package client_srv_id

import (
	"fmt"

	github_com_eden_framework_courier "github.com/eden-framework/courier"
	github_com_eden_framework_courier_client "github.com/eden-framework/courier/client"
	github_com_eden_framework_courier_status_error "github.com/eden-framework/courier/status_error"
)

type ClientSrvIDInterface interface {
	GenerateID(metas ...github_com_eden_framework_courier.Metadata) (resp *GenerateIDResponse, err error)
}

type ClientSrvID struct {
	*github_com_eden_framework_courier_client.Client
}

func (c *ClientSrvID) MarshalDefaults() {
	c.Name = "srv-id"
	c.Client.MarshalDefaults(c.Client)
}

func (c *ClientSrvID) Init() {
	c.MarshalDefaults()
	c.CheckService()
}

func (c ClientSrvID) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := github_com_eden_framework_courier_status_error.FromError(err)
	if statusErr.Code == int64(github_com_eden_framework_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

func (c ClientSrvID) GenerateID(metas ...github_com_eden_framework_courier.Metadata) (resp *GenerateIDResponse, err error) {
	resp = &GenerateIDResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".GenerateID", "GET", "/id/v0/id", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GenerateIDResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body GenerateIDResp
}
