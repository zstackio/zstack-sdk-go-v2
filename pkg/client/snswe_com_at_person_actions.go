// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSWeComAtPerson 查询SNSWeComAtPerson列表
func (cli *ZSClient) QuerySNSWeComAtPerson(params param.QueryParam) ([]view.QuerySNSWeComAtPersonView, error) {
	var resp []view.QuerySNSWeComAtPersonView
	return resp, cli.List("v1/sns/application-endpoints/we-com/at-persons", &params, &resp)
}

