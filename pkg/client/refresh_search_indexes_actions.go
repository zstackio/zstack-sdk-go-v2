// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshSearchIndexes operates on RefreshSearchIndexes
func (cli *ZSClient) RefreshSearchIndexes(params param.RefreshSearchIndexesParam) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.Get("v1/search/indexes/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
