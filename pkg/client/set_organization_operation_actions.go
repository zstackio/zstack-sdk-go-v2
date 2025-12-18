// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetOrganizationOperation 操作SetOrganizationOperation
func (cli *ZSClient) SetOrganizationOperation(uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/operation", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

