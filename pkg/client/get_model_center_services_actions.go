// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetModelCenterServices gets ModelCenterServices by uuid
func (cli *ZSClient) GetModelCenterServices(uuid string) (*view.GetModelCenterServicesView, error) {
	var resp view.GetModelCenterServicesView
	if err := cli.Get("v1/ai/model-centers/services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
