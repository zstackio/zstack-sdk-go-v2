// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(uuid string, params param.UpdateDatasetParam) (*view.UpdateDatasetEventView, error) {
	resp := view.UpdateDatasetEventView{}
	if err := cli.Put("v1/ai/datasets/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
