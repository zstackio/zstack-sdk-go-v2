// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIAM2VirtualIDType changes IAM2VirtualIDType
func (cli *ZSClient) ChangeIAM2VirtualIDType(uuid string, params param.ChangeIAM2VirtualIDTypeParam) (*view.ChangeIAM2VirtualIDTypeEventView, error) {
	resp := view.ChangeIAM2VirtualIDTypeEventView{}
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
