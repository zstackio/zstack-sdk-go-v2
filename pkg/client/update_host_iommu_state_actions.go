// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateHostIommuState updates HostIommuState
func (cli *ZSClient) UpdateHostIommuState(uuid string, params param.UpdateHostIommuStateParam) (*view.UpdateHostIommuStateEventView, error) {
	resp := view.UpdateHostIommuStateEventView{}
	if err := cli.Put("v1/pci-device/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
