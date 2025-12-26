// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIAM2VirtualID updates IAM2VirtualID
func (cli *ZSClient) UpdateIAM2VirtualID(uuid string, params param.UpdateIAM2VirtualIDParam) (*view.UpdateIAM2VirtualIDEventView, error) {
	resp := view.UpdateIAM2VirtualIDEventView{}
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
