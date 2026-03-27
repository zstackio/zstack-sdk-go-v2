// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSEmailPlatform creates SNSEmailPlatform
func (cli *ZSClient) CreateSNSEmailPlatform(ctx context.Context, params param.CreateSNSEmailPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-platforms/email", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSNSEmailPlatform operates on SNSEmailPlatform
func (cli *ZSClient) ValidateSNSEmailPlatform(ctx context.Context, uuid string, params param.ValidateSNSEmailPlatformParam) (*view.SNSEmailPlatformInventoryView, error) {
	resp := view.SNSEmailPlatformInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-platforms/email", uuid, "", map[string]interface{}{
		"validateSNSEmailPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSEmailPlatform queries SNSEmailPlatform list
func (cli *ZSClient) QuerySNSEmailPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List(ctx, "v1/sns/application-platforms/email", params, &resp)
}

func (cli *ZSClient) GetSNSEmailPlatform(ctx context.Context, uuid string) (*view.SNSEmailPlatformInventoryView, error) {
	var resp view.SNSEmailPlatformInventoryView
	if err := cli.Get(ctx, "v1/sns/application-platforms/email", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSEmailPlatform Pagination
func (cli *ZSClient) PageSNSEmailPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, int, error) {
	var sNSEmailPlatforms []view.SNSEmailPlatformInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-platforms/email", params, &sNSEmailPlatforms)
	return sNSEmailPlatforms, total, err
}
