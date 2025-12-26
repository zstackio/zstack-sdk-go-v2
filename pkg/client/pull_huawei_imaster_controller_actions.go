// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PullHuaweiIMasterController operates on PullHuaweiIMasterController
func (cli *ZSClient) PullHuaweiIMasterController(uuid string, params param.PullHuaweiIMasterControllerParam) (*view.PullHuaweiIMasterControllerEventView, error) {
	resp := view.PullHuaweiIMasterControllerEventView{}
	if err := cli.Put("v1/sdn-controller/huawei-imaster/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
