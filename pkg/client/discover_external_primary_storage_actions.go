// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DiscoverExternalPrimaryStorage operates on DiscoverExternalPrimaryStorage
func (cli *ZSClient) DiscoverExternalPrimaryStorage(params param.DiscoverExternalPrimaryStorageParam) (*view.DiscoverExternalPrimaryStorageEventView, error) {
	resp := view.DiscoverExternalPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/addon/discover", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
