// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateNfvInstProvisionConfig updates NfvInstProvisionConfig
func (cli *ZSClient) UpdateNfvInstProvisionConfig(uuid string, params param.UpdateNfvInstProvisionConfigParam) (*view.UpdateNfvInstProvisionConfigEventView, error) {
	resp := view.UpdateNfvInstProvisionConfigEventView{}
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/provision/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
