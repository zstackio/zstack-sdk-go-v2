// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BootstrapMiniHost operates on BootstrapMiniHost
func (cli *ZSClient) BootstrapMiniHost(params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.Post("v1/mini-clusters/hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
