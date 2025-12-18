// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshGuestOsMetadata 操作RefreshGuestOsMetadata
func (cli *ZSClient) RefreshGuestOsMetadata(uuid string, params param.RefreshGuestOsMetadataParam) (*view.RefreshGuestOsMetadataEventView, error) {
	resp := view.RefreshGuestOsMetadataEventView{}
	if err := cli.Put("v1/guest-os/metadata/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

