// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetEipAttachableVmNics 获取EipAttachableVmNics详情
func (cli *ZSClient) GetEipAttachableVmNics(uuid string) (*view.GetEipAttachableVmNicsView, error) {
	var resp view.GetEipAttachableVmNicsView
	if err := cli.Get("v1/eips/{eipUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

