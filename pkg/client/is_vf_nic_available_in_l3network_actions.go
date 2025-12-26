// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// IsVfNicAvailableInL3Network operates on IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(params param.IsVfNicAvailableInL3NetworkParam) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/hosts/{hostUuid}/vfnicavailable", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
