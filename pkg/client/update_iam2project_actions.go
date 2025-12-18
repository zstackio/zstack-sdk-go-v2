// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2Project updates IAM2Project
func (cli *ZSClient) UpdateIAM2Project(uuid string, params param.UpdateIAM2ProjectParam) (*view.UpdateIAM2ProjectEventView, error) {
	resp := view.UpdateIAM2ProjectEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
