// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2VirtualIDAPIPermission gets IAM2VirtualIDAPIPermission by uuid
func (cli *ZSClient) GetIAM2VirtualIDAPIPermission(uuid string) (*view.GetIAM2VirtualIDAPIPermissionView, error) {
	var resp view.GetIAM2VirtualIDAPIPermissionView
	if err := cli.Get("v1/iam2/virtual-ids/api-permissions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
