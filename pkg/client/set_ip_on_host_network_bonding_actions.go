// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetIpOnHostNetworkBonding operates on SetIpOnHostNetworkBonding
func (cli *ZSClient) SetIpOnHostNetworkBonding(params param.SetIpOnHostNetworkBondingParam) (*view.SetIpOnHostNetworkBondingEventView, error) {
	resp := view.SetIpOnHostNetworkBondingEventView{}
	if err := cli.Post("v1/hosts/bondings/{bondingUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
