// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachNfvInstToGroup operates on NfvInstToGroup
func (cli *ZSClient) AttachNfvInstToGroup(uuid string, params param.AttachNfvInstToGroupParam) (*view.AttachNfvInstToGroupEventView, error) {
	resp := view.AttachNfvInstToGroupEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{groupUuid}/instances/{nfvInstUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
