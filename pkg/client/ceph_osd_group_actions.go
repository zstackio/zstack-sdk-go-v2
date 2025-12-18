// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCephOsdGroup 查询CephOsdGroup列表
func (cli *ZSClient) QueryCephOsdGroup(params param.QueryParam) ([]view.QueryCephOsdGroupView, error) {
	var resp []view.QueryCephOsdGroupView
	return resp, cli.List("v1/primary-storage/ceph/osdgroups", &params, &resp)
}

