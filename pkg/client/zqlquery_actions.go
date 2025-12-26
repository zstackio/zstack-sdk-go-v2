// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ZQLQuery operates on ZQLQuery
func (cli *ZSClient) ZQLQuery(params param.ZQLQueryParam) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.Get("v1/zql", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
