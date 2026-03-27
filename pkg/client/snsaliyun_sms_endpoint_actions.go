// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSAliyunSmsEndpoint creates SNSAliyunSmsEndpoint
func (cli *ZSClient) CreateSNSAliyunSmsEndpoint(ctx context.Context, params param.CreateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	resp := view.SNSAliyunSmsEndpointInventoryView{}
	if err := cli.Post(ctx, "v1/sns/sms-endpoints/aliyunsms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSNSAliyunSmsEndpoint operates on SNSAliyunSmsEndpoint
func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(ctx context.Context, uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	resp := view.SNSAliyunSmsEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/sms-endpoints", uuid, "", map[string]interface{}{
		"validateSNSAliyunSmsEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
