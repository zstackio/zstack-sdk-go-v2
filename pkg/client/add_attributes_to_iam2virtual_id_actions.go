// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAttributesToIAM2VirtualID adds AttributesToIAM2VirtualID
func (cli *ZSClient) AddAttributesToIAM2VirtualID(params param.AddAttributesToIAM2VirtualIDParam) (*view.AddAttributesToIAM2VirtualIDEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
