// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSFeiShuAtPerson queries SNSFeiShuAtPerson list
func (cli *ZSClient) QuerySNSFeiShuAtPerson(params *param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, error) {
	var resp []view.SNSFeiShuAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu/at-persons", params, &resp)
}
// RemoveSNSFeiShuAtPerson removes SNSFeiShuAtPerson
func (cli *ZSClient) RemoveSNSFeiShuAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/feishu/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}
// AddSNSFeiShuAtPerson adds SNSFeiShuAtPerson
func (cli *ZSClient) AddSNSFeiShuAtPerson(params param.AddSNSFeiShuAtPersonParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	var resp view.AddSNSFeiShuAtPersonEventView
	if err := cli.Post("v1/sns/application-endpoints/feishu/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
