// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSecurityGroupState changes SecurityGroupState
func (cli *ZSClient) ChangeSecurityGroupState(uuid string, params param.ChangeSecurityGroupStateParam) (*view.ChangeSecurityGroupStateEventView, error) {
	resp := view.ChangeSecurityGroupStateEventView{}
	if err := cli.Put("v1/security-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
