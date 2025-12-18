// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSDingTalkAtPerson queries SNSDingTalkAtPerson list
func (cli *ZSClient) QuerySNSDingTalkAtPerson(params param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, error) {
	var resp []view.SNSDingTalkAtPersonInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk/at-persons", &params, &resp)
}
