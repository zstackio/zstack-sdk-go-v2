// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachNicToBonding 操作NicToBonding
func (cli *ZSClient) AttachNicToBonding(uuid string, params param.AttachNicToBondingParam) (*view.AttachNicToBondingEventView, error) {
	resp := view.AttachNicToBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

