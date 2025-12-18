// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVipQos 操作SetVipQos
func (cli *ZSClient) SetVipQos(uuid string, params param.SetVipQosParam) (*view.SetVipQosEventView, error) {
	resp := view.SetVipQosEventView{}
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

