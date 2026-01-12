// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveSNSDingTalkAtPerson removes SNSDingTalkAtPerson
func (cli *ZSClient) RemoveSNSDingTalkAtPerson(endpointUuid string, phoneNumber string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/application-endpoints/ding-talk", fmt.Sprintf(\"%s/at-persons/%s\", endpointUuid, phoneNumber), string(deleteMode))
}
// QuerySNSDingTalkAtPerson queries SNSDingTalkAtPerson list
func (cli *ZSClient) QuerySNSDingTalkAtPerson(params *param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, error) {
	var resp []view.SNSDingTalkAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSDingTalkAtPerson(uuid string) (*view.SNSDingTalkAtPersonInventoryView, error) {
	var resp view.SNSDingTalkAtPersonInventoryView
	if err := cli.Get("v1/sns/application-endpoints/ding-talk/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSNSDingTalkAtPerson adds SNSDingTalkAtPerson
func (cli *ZSClient) AddSNSDingTalkAtPerson(params param.AddSNSDingTalkAtPersonParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	var resp view.AddSNSDingTalkAtPersonEventView
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
