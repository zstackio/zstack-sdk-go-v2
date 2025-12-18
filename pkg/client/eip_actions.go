// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEip 更新Eip
func (cli *ZSClient) UpdateEip(uuid string, params param.UpdateEipParam) (*view.UpdateEipEventView, error) {
	resp := view.UpdateEipEventView{}
	if err := cli.Put("v1/eips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

