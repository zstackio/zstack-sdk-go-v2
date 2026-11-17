// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// RefreshSearchIndexes operates on SearchIndexes
func (cli *ZSClient) RefreshSearchIndexes() (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.GetWithRespKey("v1/search/indexes/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
