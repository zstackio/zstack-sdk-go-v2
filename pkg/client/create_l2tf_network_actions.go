// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2TfNetwork creates L2TfNetwork
func (cli *ZSClient) CreateL2TfNetwork(params param.CreateL2TfNetworkParam) (*view.CreateL2NetworkEventView, error) {
	resp := view.CreateL2NetworkEventView{}
	if err := cli.Post("v1/l2-networks/tf", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
