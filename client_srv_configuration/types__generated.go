package client_srv_configuration

import (
	github_com_eden_framework_courier_httpx "github.com/eden-framework/courier/httpx"
	github_com_eden_framework_courier_status_error "github.com/eden-framework/courier/status_error"
	github_com_eden_framework_courier_swagger "github.com/eden-framework/courier/swagger"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

type Configuration struct {
	//
	GithubComEdenFrameworkSqlxDatatypesPrimaryID
	//
	GithubComEdenFrameworkSqlxDatatypesOperateTime
	// 业务ID
	ConfigurationID uint64 `json:"configurationID,string"`
	// Key
	Key string `json:"key"`
	// Remark
	Remark string `json:"remark"`
	// StackID
	StackID uint64 `json:"stackID,string"`
	// Value
	Value string `json:"value"`
}

type CreateConfigurationBody struct {
	// Key
	Key string `json:"key"`
	// Remark
	Remark string `json:"remark"`
	// StackID
	StackID uint64 `json:"stackID,string"`
	// Value
	Value string `json:"value"`
}

type GithubComEdenFrameworkCourierHttpxHTML = github_com_eden_framework_courier_httpx.HTML

type GithubComEdenFrameworkCourierStatusErrorErrorField = github_com_eden_framework_courier_status_error.ErrorField

type GithubComEdenFrameworkCourierStatusErrorErrorFields = github_com_eden_framework_courier_status_error.ErrorFields

type GithubComEdenFrameworkCourierStatusErrorStatusError = github_com_eden_framework_courier_status_error.StatusError

type GithubComEdenFrameworkCourierSwaggerJSONBytes = github_com_eden_framework_courier_swagger.JSONBytes

type GithubComEdenFrameworkSqlxDatatypesOperateTime = github_com_eden_framework_sqlx_datatypes.OperateTime

type GithubComEdenFrameworkSqlxDatatypesPrimaryID = github_com_eden_framework_sqlx_datatypes.PrimaryID

type GithubComEdenFrameworkSqlxDatatypesTimestamp = github_com_eden_framework_sqlx_datatypes.Timestamp

type UpdateConfigurationBody struct {
	// Key
	Key string `json:"key"`
	// Remark
	Remark string `json:"remark"`
	// StackID
	StackID uint64 `json:"stackID,string"`
	// Value
	Value string `json:"value"`
}
