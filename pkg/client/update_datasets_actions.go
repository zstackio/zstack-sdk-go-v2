// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateDatasets updates Datasets
func (cli *ZSClient) UpdateDatasets(uuid string, params param.UpdateDatasetsParam) (*view.UpdateDatasetsEventView, error) {
	resp := view.UpdateDatasetsEventView{}
	if err := cli.Put("v1/ai/datasets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
