// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetInterdependentL3NetworksImages gets InterdependentL3NetworksImages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksImages(uuid string) (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.Get("v1/images-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
