// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmNicAttachableEips 获取VmNicAttachableEips详情
func (cli *ZSClient) GetVmNicAttachableEips(uuid string) (*view.GetVmNicAttachableEipsView, error) {
	var resp view.GetVmNicAttachableEipsView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/candidate-eips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

