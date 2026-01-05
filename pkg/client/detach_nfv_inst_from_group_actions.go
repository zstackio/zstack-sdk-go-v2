// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachNfvInstFromGroup operates on NfvInstFromGroup
func (cli *ZSClient) DetachNfvInstFromGroup(uuid string, params param.DetachNfvInstFromGroupParam) (*view.DetachNfvInstFromGroupEventView, error) {
	resp := view.DetachNfvInstFromGroupEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{groupUuid}/instances/{nfvInstUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
