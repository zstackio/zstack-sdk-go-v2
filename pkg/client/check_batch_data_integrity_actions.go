// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckBatchDataIntegrity 操作CheckBatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity(params param.CheckBatchDataIntegrityParam) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.Get("v1/check/batch/data/integrity/", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

