// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBareMetal2InstancePassword changes BareMetal2InstancePassword
func (cli *ZSClient) ChangeBareMetal2InstancePassword(uuid string, params param.ChangeBareMetal2InstancePasswordParam) (*view.ChangeBareMetal2InstancePasswordEventView, error) {
	resp := view.ChangeBareMetal2InstancePasswordEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
