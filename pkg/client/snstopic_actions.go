// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSTopic deletes SNSTopic
func (cli *ZSClient) DeleteSNSTopic(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sns/topics", uuid, string(deleteMode))
}
// UpdateSNSTopic updates SNSTopic
func (cli *ZSClient) UpdateSNSTopic(ctx context.Context, uuid string, params param.UpdateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/topics", uuid, "", map[string]interface{}{
		"updateSNSTopic": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSTopic queries SNSTopic list
func (cli *ZSClient) QuerySNSTopic(ctx context.Context, params *param.QueryParam) ([]view.SNSTopicInventoryView, error) {
	var resp []view.SNSTopicInventoryView
	return resp, cli.List(ctx, "v1/sns/topics", params, &resp)
}

func (cli *ZSClient) GetSNSTopic(ctx context.Context, uuid string) (*view.SNSTopicInventoryView, error) {
	var resp view.SNSTopicInventoryView
	if err := cli.Get(ctx, "v1/sns/topics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTopic Pagination
func (cli *ZSClient) PageSNSTopic(ctx context.Context, params *param.QueryParam) ([]view.SNSTopicInventoryView, int, error) {
	var sNSTopics []view.SNSTopicInventoryView
	total, err := cli.Page(ctx, "v1/sns/topics", params, &sNSTopics)
	return sNSTopics, total, err
}
// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(ctx context.Context, params param.CreateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.Post(ctx, "v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
