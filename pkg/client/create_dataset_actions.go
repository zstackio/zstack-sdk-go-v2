// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(params param.CreateDatasetParam) (*view.CreateDatasetEventView, error) {
	resp := view.CreateDatasetEventView{}
	if err := cli.Post("v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
