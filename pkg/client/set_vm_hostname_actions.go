// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmHostname 操作SetVmHostname
func (cli *ZSClient) SetVmHostname(uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

