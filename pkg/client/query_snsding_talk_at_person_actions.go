// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSDingTalkAtPerson queries SNSDingTalkAtPerson list
func (cli *ZSClient) QuerySNSDingTalkAtPerson(params *param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, error) {
	var resp []view.SNSDingTalkAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk/at-persons", params, &resp)
}
