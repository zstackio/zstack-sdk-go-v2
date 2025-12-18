// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAttributesToIAM2Project adds AttributesToIAM2Project
func (cli *ZSClient) AddAttributesToIAM2Project(params param.AddAttributesToIAM2ProjectParam) (*view.AddAttributesToIAM2ProjectEventView, error) {
	resp := view.AddAttributesToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
