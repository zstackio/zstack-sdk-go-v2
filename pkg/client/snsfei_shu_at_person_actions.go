// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSFeiShuAtPerson queries SNSFeiShuAtPerson list
func (cli *ZSClient) QuerySNSFeiShuAtPerson(params *param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, error) {
	var resp []view.SNSFeiShuAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSFeiShuAtPerson(uuid string) (*view.SNSFeiShuAtPersonInventoryView, error) {
	var resp view.SNSFeiShuAtPersonInventoryView
	if err := cli.Get("v1/sns/application-endpoints/feishu/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RemoveSNSFeiShuAtPerson removes SNSFeiShuAtPerson
func (cli *ZSClient) RemoveSNSFeiShuAtPerson(endpointUuid string, userId string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/application-endpoints/feishu", endpointUuid, fmt.Sprintf("at-persons/%s", userId), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddSNSFeiShuAtPerson adds SNSFeiShuAtPerson
func (cli *ZSClient) AddSNSFeiShuAtPerson(params param.AddSNSFeiShuAtPersonParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	var resp view.AddSNSFeiShuAtPersonEventView
	if err := cli.Post("v1/sns/application-endpoints/feishu/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
