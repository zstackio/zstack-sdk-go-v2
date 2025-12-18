// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PullHuaweiIMasterController 操作PullHuaweiIMasterController
func (cli *ZSClient) PullHuaweiIMasterController(uuid string, params param.PullHuaweiIMasterControllerParam) (*view.PullHuaweiIMasterControllerEventView, error) {
	resp := view.PullHuaweiIMasterControllerEventView{}
	if err := cli.Put("v1/sdn-controller/huawei-imaster/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

