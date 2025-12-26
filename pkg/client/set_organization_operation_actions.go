// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetOrganizationOperation operates on SetOrganizationOperation
func (cli *ZSClient) SetOrganizationOperation(uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/operation", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
