// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddResourceToIAM2Project adds ResourceToIAM2Project
func (cli *ZSClient) AddResourceToIAM2Project(params param.AddResourceToIAM2ProjectParam) (*view.AddResourceToIAM2ProjectEventView, error) {
	resp := view.AddResourceToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/add/resource/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
