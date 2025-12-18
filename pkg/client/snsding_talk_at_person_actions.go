// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSDingTalkAtPerson 查询SNSDingTalkAtPerson列表
func (cli *ZSClient) QuerySNSDingTalkAtPerson(params param.QueryParam) ([]view.QuerySNSDingTalkAtPersonView, error) {
	var resp []view.QuerySNSDingTalkAtPersonView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk/at-persons", &params, &resp)
}

