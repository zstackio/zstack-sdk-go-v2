// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmNicAttachedNetworkService 获取VmNicAttachedNetworkService详情
func (cli *ZSClient) GetVmNicAttachedNetworkService(uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/attached-networkservices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

