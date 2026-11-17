// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetGuestOsMetadata gets GuestOsMetadata by uuid
func (cli *ZSClient) GetGuestOsMetadata() (*view.GuestOsCharacterInventoryView, error) {
	var resp view.GuestOsCharacterInventoryView
	if err := cli.GetWithRespKey("v1/guest-os/metadata", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshGuestOsMetadata operates on GuestOsMetadata
func (cli *ZSClient) RefreshGuestOsMetadata() (*view.RefreshGuestOsMetadataEventView, error) {
	resp := view.RefreshGuestOsMetadataEventView{}
	if err := cli.PutWithRespKey("v1/guest-os/metadata/actions", "", "", map[string]interface{}{
		"refreshGuestOsMetadata": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
