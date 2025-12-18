// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(uuid string, params param.UpdateDatasetParam) (*view.UpdateDatasetEventView, error) {
	resp := view.UpdateDatasetEventView{}
	if err := cli.Put("v1/ai/datasets/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
