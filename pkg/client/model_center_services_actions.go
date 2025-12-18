// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetModelCenterServices 获取ModelCenterServices详情
func (cli *ZSClient) GetModelCenterServices(uuid string) (*view.GetModelCenterServicesView, error) {
	var resp view.GetModelCenterServicesView
	if err := cli.Get("v1/ai/model-centers/services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

