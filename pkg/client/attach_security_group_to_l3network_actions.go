// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachSecurityGroupToL3Network operates on SecurityGroupToL3Network
func (cli *ZSClient) AttachSecurityGroupToL3Network(params param.AttachSecurityGroupToL3NetworkParam) (*view.AttachSecurityGroupToL3NetworkEventView, error) {
	resp := view.AttachSecurityGroupToL3NetworkEventView{}
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
