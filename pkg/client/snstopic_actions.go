// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSTopic deletes SNSTopic
func (cli *ZSClient) DeleteSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics", uuid, string(deleteMode))
}
// UpdateSNSTopic updates SNSTopic
func (cli *ZSClient) UpdateSNSTopic(uuid string, params param.UpdateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	var resp view.UpdateSNSTopicEventView
	if err := cli.Put("v1/sns/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSTopic queries SNSTopic list
func (cli *ZSClient) QuerySNSTopic(params *param.QueryParam) ([]view.SNSTopicInventoryView, error) {
	var resp []view.SNSTopicInventoryView
	return resp, cli.List("v1/sns/topics", params, &resp)
}

func (cli *ZSClient) GetSNSTopic(uuid string) (*view.SNSTopicInventoryView, error) {
	var resp view.SNSTopicInventoryView
	if err := cli.Get("v1/sns/topics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(params param.CreateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	var resp view.CreateSNSTopicEventView
	if err := cli.Post("v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
