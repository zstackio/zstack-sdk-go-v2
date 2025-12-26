// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHypervisorTypes gets HypervisorTypes by uuid
func (cli *ZSClient) GetHypervisorTypes(uuid string) (*view.GetHypervisorTypesView, error) {
	var resp view.GetHypervisorTypesView
	if err := cli.Get("v1/hosts/hypervisor-types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
