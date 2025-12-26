// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAllEventMetadata gets AllEventMetadata by uuid
func (cli *ZSClient) GetAllEventMetadata(uuid string) (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.Get("v1/zwatch/events/meta-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
