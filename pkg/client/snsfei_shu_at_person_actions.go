// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSFeiShuAtPerson 查询SNSFeiShuAtPerson列表
func (cli *ZSClient) QuerySNSFeiShuAtPerson(params param.QueryParam) ([]view.QuerySNSFeiShuAtPersonView, error) {
	var resp []view.QuerySNSFeiShuAtPersonView
	return resp, cli.List("v1/sns/application-endpoints/feishu/at-persons", &params, &resp)
}

