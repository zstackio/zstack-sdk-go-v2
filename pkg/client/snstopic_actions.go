// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSTopic deletes SNSTopic
func (cli *ZSClient) DeleteSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{uuid}", uuid, string(deleteMode))
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
// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(params param.CreateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	var resp view.CreateSNSTopicEventView
	if err := cli.Post("v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
