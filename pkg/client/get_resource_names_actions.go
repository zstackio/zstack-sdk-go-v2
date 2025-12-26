// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames(uuid string) (*view.GetResourceNamesView, error) {
	var resp view.GetResourceNamesView
	if err := cli.Get("v1/resources/names", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
