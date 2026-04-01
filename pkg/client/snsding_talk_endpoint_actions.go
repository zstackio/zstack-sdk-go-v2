// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSDingTalkEndpoint creates SNSDingTalkEndpoint
func (cli *ZSClient) CreateSNSDingTalkEndpoint(params param.CreateSNSDingTalkEndpointParam) (*view.SNSDingTalkEndpointInventoryView, error) {
	resp := view.SNSDingTalkEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSDingTalkEndpoint queries SNSDingTalkEndpoint list
func (cli *ZSClient) QuerySNSDingTalkEndpoint(params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, error) {
	var resp []view.SNSDingTalkEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk", params, &resp)
}

func (cli *ZSClient) GetSNSDingTalkEndpoint(uuid string) (*view.SNSDingTalkEndpointInventoryView, error) {
	var resp view.SNSDingTalkEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/ding-talk", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSDingTalkEndpoint Pagination
func (cli *ZSClient) PageSNSDingTalkEndpoint(params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, int, error) {
	var sNSDingTalkEndpoints []view.SNSDingTalkEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/ding-talk", params, &sNSDingTalkEndpoints)
	return sNSDingTalkEndpoints, total, err
}
// UpdateSNSDingTalkEndpoint updates SNSDingTalkEndpoint
func (cli *ZSClient) UpdateSNSDingTalkEndpoint(uuid string, params param.UpdateSNSDingTalkEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/application-endpoints/ding-talk", uuid, "", map[string]interface{}{
		"updateSNSDingTalkEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
