// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeHostPassword 操作HostPassword
func (cli *ZSClient) ChangeHostPassword(uuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.Put("v1/hosts/kvm/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

