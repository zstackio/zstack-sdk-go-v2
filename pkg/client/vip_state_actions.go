// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVipState 操作VipState
func (cli *ZSClient) ChangeVipState(uuid string, params param.ChangeVipStateParam) (*view.ChangeVipStateEventView, error) {
	resp := view.ChangeVipStateEventView{}
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

