// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BatchQuery operates on BatchQuery
func (cli *ZSClient) BatchQuery(params param.BatchQueryParam) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.Get("v1/batch-queries", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
