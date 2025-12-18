// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BatchQuery operates on BatchQuery
func (cli *ZSClient) BatchQuery(params param.BatchQueryParam) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.Get("v1/batch-queries", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
