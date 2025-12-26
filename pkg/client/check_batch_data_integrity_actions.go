// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckBatchDataIntegrity operates on CheckBatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity(params param.CheckBatchDataIntegrityParam) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.Get("v1/check/batch/data/integrity/", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
