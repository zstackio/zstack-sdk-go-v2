// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAttributesToIAM2Project adds AttributesToIAM2Project
func (cli *ZSClient) AddAttributesToIAM2Project(params param.AddAttributesToIAM2ProjectParam) (*view.AddAttributesToIAM2ProjectEventView, error) {
	resp := view.AddAttributesToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
