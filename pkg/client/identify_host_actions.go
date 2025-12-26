// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// IdentifyHost operates on IdentifyHost
func (cli *ZSClient) IdentifyHost(uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
