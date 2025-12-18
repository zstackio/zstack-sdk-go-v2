// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverIAM2Project operates on IAM2Project
func (cli *ZSClient) RecoverIAM2Project(uuid string, params param.RecoverIAM2ProjectParam) (*view.RecoverIAM2ProjectEventView, error) {
	resp := view.RecoverIAM2ProjectEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
