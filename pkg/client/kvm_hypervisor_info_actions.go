// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryKvmHypervisorInfo 查询KvmHypervisorInfo列表
func (cli *ZSClient) QueryKvmHypervisorInfo(params param.QueryParam) ([]view.QueryKvmHypervisorInfoView, error) {
	var resp []view.QueryKvmHypervisorInfoView
	return resp, cli.List("v1/hosts/kvm/hypervisor/info", &params, &resp)
}

