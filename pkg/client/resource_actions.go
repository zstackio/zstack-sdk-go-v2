// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames() (*view.ResourceInventoryView, error) {
	var resp view.ResourceInventoryView
	if err := cli.GetWithRespKey("v1/resources/names", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
