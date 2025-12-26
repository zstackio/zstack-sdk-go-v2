// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetNetworkServiceTypes gets NetworkServiceTypes by uuid
func (cli *ZSClient) GetNetworkServiceTypes(uuid string) (*view.GetNetworkServiceTypesView, error) {
	var resp view.GetNetworkServiceTypesView
	if err := cli.Get("v1/network-services/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
