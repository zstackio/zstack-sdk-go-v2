// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateDatasets updates Datasets
func (cli *ZSClient) UpdateDatasets(uuid string, params param.UpdateDatasetsParam) (*view.UpdateDatasetsEventView, error) {
	resp := view.UpdateDatasetsEventView{}
	if err := cli.Put("v1/ai/datasets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
