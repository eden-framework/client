package client_srv_configuration

import (
	"fmt"

	github_com_eden_framework_courier "github.com/eden-framework/courier"
	github_com_eden_framework_courier_client "github.com/eden-framework/courier/client"
	github_com_eden_framework_courier_status_error "github.com/eden-framework/courier/status_error"
)

type ClientSrvConfigurationInterface interface {
	BatchCreateConfigs(req BatchCreateConfigsRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *BatchCreateConfigsResponse, err error)
	CreateConfig(req CreateConfigRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *CreateConfigResponse, err error)
	GetConfigurations(req GetConfigurationsRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *GetConfigurationsResponse, err error)
	SwaggerDoc(metas ...github_com_eden_framework_courier.Metadata) (resp *SwaggerDocResponse, err error)
	SwaggerUIBundle(metas ...github_com_eden_framework_courier.Metadata) (resp *SwaggerUIBundleResponse, err error)
	UpdateConfig(req UpdateConfigRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *UpdateConfigResponse, err error)
}

type ClientSrvConfiguration struct {
	*github_com_eden_framework_courier_client.Client
}

func (c *ClientSrvConfiguration) MarshalDefaults() {
	c.Name = "srv-configuration"
	c.Client.MarshalDefaults(c.Client)
}

func (c *ClientSrvConfiguration) Init() {
	c.MarshalDefaults()
	c.CheckService()
}

func (c ClientSrvConfiguration) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := github_com_eden_framework_courier_status_error.FromError(err)
	if statusErr.Code == int64(github_com_eden_framework_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

type BatchCreateConfigsRequest struct {
	//
	Body []CreateConfigurationBody `fmt:"json" in:"body"`
}

func (c ClientSrvConfiguration) BatchCreateConfigs(req BatchCreateConfigsRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *BatchCreateConfigsResponse, err error) {
	resp = &BatchCreateConfigsResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".BatchCreateConfigs", "POST", "/configuration/v0/configs/batch", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type BatchCreateConfigsResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []byte
}

type CreateConfigRequest struct {
	//
	Body struct {
		//
		CreateConfigurationBody
	} `fmt:"json" in:"body"`
}

func (c ClientSrvConfiguration) CreateConfig(req CreateConfigRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *CreateConfigResponse, err error) {
	resp = &CreateConfigResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".CreateConfig", "POST", "/configuration/v0/configs", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateConfigResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []byte
}

type GetConfigurationsRequest struct {
	//
	StackID uint64 `in:"query" name:"stackID"`
}

func (c ClientSrvConfiguration) GetConfigurations(req GetConfigurationsRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *GetConfigurationsResponse, err error) {
	resp = &GetConfigurationsResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".GetConfigurations", "GET", "/configuration/v0/configs", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetConfigurationsResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []Configuration
}

func (c ClientSrvConfiguration) SwaggerDoc(metas ...github_com_eden_framework_courier.Metadata) (resp *SwaggerDocResponse, err error) {
	resp = &SwaggerDocResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".SwaggerDoc", "GET", "/configuration/doc", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerDocResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []byte
}

func (c ClientSrvConfiguration) SwaggerUIBundle(metas ...github_com_eden_framework_courier.Metadata) (resp *SwaggerUIBundleResponse, err error) {
	resp = &SwaggerUIBundleResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".SwaggerUIBundle", "GET", "/configuration/swagger-ui-bundle.js", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerUIBundleResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []byte
}

type UpdateConfigRequest struct {
	//
	ConfigID uint64 `in:"path" name:"configID"`
	//
	Body struct {
		//
		UpdateConfigurationBody
	} `fmt:"json" in:"body"`
}

func (c ClientSrvConfiguration) UpdateConfig(req UpdateConfigRequest, metas ...github_com_eden_framework_courier.Metadata) (resp *UpdateConfigResponse, err error) {
	resp = &UpdateConfigResponse{}
	resp.Meta = github_com_eden_framework_courier.Metadata{}

	err = c.Request(c.Name+".UpdateConfig", "PATCH", "/configuration/v0/configs/:configID", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UpdateConfigResponse struct {
	Meta github_com_eden_framework_courier.Metadata
	Body []byte
}
