// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeHostPassword changes HostPassword
func (cli *ZSClient) ChangeHostPassword(uuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.Put("v1/hosts/kvm/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
