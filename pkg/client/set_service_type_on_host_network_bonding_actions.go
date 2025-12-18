// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetServiceTypeOnHostNetworkBonding 操作SetServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(params param.SetServiceTypeOnHostNetworkBondingParam) (*view.SetServiceTypeOnHostNetworkBondingEventView, error) {
	resp := view.SetServiceTypeOnHostNetworkBondingEventView{}
	if err := cli.Post("v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

