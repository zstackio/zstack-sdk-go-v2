// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostNetworkFacts gets HostNetworkFacts by uuid
func (cli *ZSClient) GetHostNetworkFacts(uuid string) (*view.GetHostNetworkFactsView, error) {
	var resp view.GetHostNetworkFactsView
	if err := cli.Get("v1/hosts/network-facts/{hostUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
