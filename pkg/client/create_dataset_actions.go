// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(params param.CreateDatasetParam) (*view.CreateDatasetEventView, error) {
	resp := view.CreateDatasetEventView{}
	if err := cli.Post("v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
