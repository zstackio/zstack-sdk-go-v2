// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAudit 查询Audit列表
func (cli *ZSClient) QueryAudit(params param.QueryParam) ([]view.QueryAuditView, error) {
	var resp []view.QueryAuditView
	return resp, cli.List("v1/zwatch/audit-records", &params, &resp)
}

