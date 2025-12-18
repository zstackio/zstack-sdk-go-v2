// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateExternalPrimaryStorage updates ExternalPrimaryStorage
func (cli *ZSClient) UpdateExternalPrimaryStorage(uuid string, params param.UpdateExternalPrimaryStorageParam) (*view.UpdateExternalPrimaryStorageEventView, error) {
	resp := view.UpdateExternalPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/addon/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
