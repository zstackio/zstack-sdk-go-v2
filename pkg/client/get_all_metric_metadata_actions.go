// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAllMetricMetadata gets AllMetricMetadata by uuid
func (cli *ZSClient) GetAllMetricMetadata(uuid string) (*view.GetAllMetricMetadataView, error) {
	var resp view.GetAllMetricMetadataView
	if err := cli.Get("v1/zwatch/metrics/meta-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
