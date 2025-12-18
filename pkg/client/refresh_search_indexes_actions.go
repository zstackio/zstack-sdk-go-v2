// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshSearchIndexes 操作RefreshSearchIndexes
func (cli *ZSClient) RefreshSearchIndexes(params param.RefreshSearchIndexesParam) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.Get("v1/search/indexes/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

