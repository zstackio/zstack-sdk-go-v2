// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// BatchQuery operates on Query
func (cli *ZSClient) BatchQuery() (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.GetWithRespKey("v1/batch-queries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZQLQuery operates on ZQLQuery
func (cli *ZSClient) ZQLQuery() (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.GetWithRespKey("v1/zql", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
