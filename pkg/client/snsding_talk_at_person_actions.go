// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveSNSDingTalkAtPerson removes SNSDingTalkAtPerson
func (cli *ZSClient) RemoveSNSDingTalkAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/ding-talk/{endpointUuid}/at-persons/{phoneNumber}", uuid, string(deleteMode))
}
// QuerySNSDingTalkAtPerson queries SNSDingTalkAtPerson list
func (cli *ZSClient) QuerySNSDingTalkAtPerson(params *param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, error) {
	var resp []view.SNSDingTalkAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp)
}
// AddSNSDingTalkAtPerson adds SNSDingTalkAtPerson
func (cli *ZSClient) AddSNSDingTalkAtPerson(params param.AddSNSDingTalkAtPersonParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	var resp view.AddSNSDingTalkAtPersonEventView
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
