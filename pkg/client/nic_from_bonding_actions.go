// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachNicFromBonding 操作NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(uuid string, params param.DetachNicFromBondingParam) (*view.DetachNicFromBondingEventView, error) {
	resp := view.DetachNicFromBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

