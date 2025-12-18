// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryQuota 查询Quota列表
func (cli *ZSClient) QueryQuota(params param.QueryParam) ([]view.QueryQuotaView, error) {
	var resp []view.QueryQuotaView
	return resp, cli.List("v1/accounts/quotas", &params, &resp)
}

