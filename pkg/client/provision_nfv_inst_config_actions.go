// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ProvisionNfvInstConfig operates on ProvisionNfvInstConfig
func (cli *ZSClient) ProvisionNfvInstConfig(uuid string, params param.ProvisionNfvInstConfigParam) (*view.ProvisionNfvInstConfigEventView, error) {
	resp := view.ProvisionNfvInstConfigEventView{}
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
