// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateHostIpmi 更新HostIpmi
func (cli *ZSClient) UpdateHostIpmi(uuid string, params param.UpdateHostIpmiParam) (*view.UpdateHostIpmiEventView, error) {
	resp := view.UpdateHostIpmiEventView{}
	if err := cli.Put("v1/hosts/ipmi/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

