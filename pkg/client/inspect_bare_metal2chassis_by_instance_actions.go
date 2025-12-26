// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// InspectBareMetal2ChassisByInstance operates on InspectBareMetal2ChassisByInstance
func (cli *ZSClient) InspectBareMetal2ChassisByInstance(uuid string, params param.InspectBareMetal2ChassisByInstanceParam) (*view.InspectBareMetal2ChassisByInstanceEventView, error) {
	resp := view.InspectBareMetal2ChassisByInstanceEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
