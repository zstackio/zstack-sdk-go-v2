// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectBareMetal2Instance 操作ReconnectBareMetal2Instance
func (cli *ZSClient) ReconnectBareMetal2Instance(uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.ReconnectBareMetal2InstanceEventView, error) {
	resp := view.ReconnectBareMetal2InstanceEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

