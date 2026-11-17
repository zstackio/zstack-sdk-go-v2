// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSAliyunSmsEndpoint creates SNSAliyunSmsEndpoint
func (cli *ZSClient) CreateSNSAliyunSmsEndpoint(params param.CreateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	resp := view.SNSAliyunSmsEndpointInventoryView{}
	if err := cli.Post("v1/sns/sms-endpoints/aliyunsms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSNSAliyunSmsEndpoint operates on SNSAliyunSmsEndpoint
func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	resp := view.SNSAliyunSmsEndpointInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/sms-endpoints", uuid, "", map[string]interface{}{
		"validateSNSAliyunSmsEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp []view.SNSAliyunSmsEndpointInventoryView
	return resp, cli.List("v1/sns/sms-endpoints", params, &resp)
}

func (cli *ZSClient) GetSNSSmsEndpoint(uuid string) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp view.SNSAliyunSmsEndpointInventoryView
	if err := cli.Get("v1/sns/sms-endpoints", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSSmsEndpoint Pagination
func (cli *ZSClient) PageSNSSmsEndpoint(params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, int, error) {
	var sNSSmsEndpoints []view.SNSAliyunSmsEndpointInventoryView
	total, err := cli.Page("v1/sns/sms-endpoints", params, &sNSSmsEndpoints)
	return sNSSmsEndpoints, total, err
}
