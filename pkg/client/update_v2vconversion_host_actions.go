// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateV2VConversionHost updates V2VConversionHost
func (cli *ZSClient) UpdateV2VConversionHost(uuid string, params param.UpdateV2VConversionHostParam) (*view.UpdateV2VConversionHostEventView, error) {
	resp := view.UpdateV2VConversionHostEventView{}
	if err := cli.Put("v1/v2v-conversion-hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
