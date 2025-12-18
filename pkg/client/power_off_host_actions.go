// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerOffHost 操作PowerOffHost
func (cli *ZSClient) PowerOffHost(uuid string, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.Put("v1/hosts/power-off/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

