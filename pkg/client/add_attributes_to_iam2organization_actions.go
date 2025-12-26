// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAttributesToIAM2Organization adds AttributesToIAM2Organization
func (cli *ZSClient) AddAttributesToIAM2Organization(params param.AddAttributesToIAM2OrganizationParam) (*view.AddAttributesToIAM2OrganizationEventView, error) {
	resp := view.AddAttributesToIAM2OrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
