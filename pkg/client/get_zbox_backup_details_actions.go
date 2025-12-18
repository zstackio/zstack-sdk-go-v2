// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetZBoxBackupDetails gets ZBoxBackupDetails by uuid
func (cli *ZSClient) GetZBoxBackupDetails(uuid string) (*view.GetZBoxBackupDetailsView, error) {
	var resp view.GetZBoxBackupDetailsView
	if err := cli.Get("v1/externalbackup/zbox/{uuid}/details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
