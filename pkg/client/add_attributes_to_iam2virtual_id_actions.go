// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAttributesToIAM2VirtualID adds AttributesToIAM2VirtualID
func (cli *ZSClient) AddAttributesToIAM2VirtualID(params param.AddAttributesToIAM2VirtualIDParam) (*view.AddAttributesToIAM2VirtualIDEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
