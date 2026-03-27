// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSUniversalSmsEndpoint queries SNSUniversalSmsEndpoint list
func (cli *ZSClient) QuerySNSUniversalSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp []view.SNSUniversalSmsEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/universal-sms", params, &resp)
}

func (cli *ZSClient) GetSNSUniversalSmsEndpoint(ctx context.Context, uuid string) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp view.SNSUniversalSmsEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/universal-sms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSUniversalSmsEndpoint Pagination
func (cli *ZSClient) PageSNSUniversalSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, int, error) {
	var sNSUniversalSmsEndpoints []view.SNSUniversalSmsEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/universal-sms", params, &sNSUniversalSmsEndpoints)
	return sNSUniversalSmsEndpoints, total, err
}
// UpdateSNSUniversalSmsEndpoint updates SNSUniversalSmsEndpoint
func (cli *ZSClient) UpdateSNSUniversalSmsEndpoint(ctx context.Context, uuid string, params param.UpdateSNSUniversalSmsEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/universal-sms", uuid, "", map[string]interface{}{
		"updateSNSUniversalSmsEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSUniversalSmsEndpoint creates SNSUniversalSmsEndpoint
func (cli *ZSClient) CreateSNSUniversalSmsEndpoint(ctx context.Context, params param.CreateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	resp := view.SNSUniversalSmsEndpointInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/universal-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSNSUniversalSmsEndpoint operates on SNSUniversalSmsEndpoint
func (cli *ZSClient) ValidateSNSUniversalSmsEndpoint(ctx context.Context, uuid string, params param.ValidateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	resp := view.SNSUniversalSmsEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/universal-sms", uuid, "", map[string]interface{}{
		"validateSNSUniversalSmsEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
