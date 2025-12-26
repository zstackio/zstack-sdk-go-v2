// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHostIpmi updates HostIpmi
func (cli *ZSClient) UpdateHostIpmi(uuid string, params param.UpdateHostIpmiParam) (*view.UpdateHostIpmiEventView, error) {
	resp := view.UpdateHostIpmiEventView{}
	if err := cli.Put("v1/hosts/ipmi/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
