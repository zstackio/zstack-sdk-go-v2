// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAccessPath gets AccessPath by uuid
func (cli *ZSClient) GetAccessPath(uuid string) (*view.GetAccessPathView, error) {
	var resp view.GetAccessPathView
	if err := cli.Get("v1/block-volumes/access/path", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
