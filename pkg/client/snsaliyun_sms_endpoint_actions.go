// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSAliyunSmsEndpoint creates SNSAliyunSmsEndpoint
func (cli *ZSClient) CreateSNSAliyunSmsEndpoint(params param.CreateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp view.CreateSNSAliyunSmsEndpointEventView
	if err := cli.Post("v1/sns/sms-endpoints/aliyunsms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ValidateSNSAliyunSmsEndpoint operates on SNSAliyunSmsEndpoint
func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	resp := view.SNSAliyunSmsEndpointInventoryView{}
	if err := cli.Put("v1/sns/sms-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
