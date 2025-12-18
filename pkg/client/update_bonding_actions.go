// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(uuid string, params param.UpdateBondingParam) (*view.UpdateBondingEventView, error) {
	resp := view.UpdateBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
