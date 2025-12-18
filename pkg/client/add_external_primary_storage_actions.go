// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddExternalPrimaryStorage adds ExternalPrimaryStorage
func (cli *ZSClient) AddExternalPrimaryStorage(params param.AddExternalPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
