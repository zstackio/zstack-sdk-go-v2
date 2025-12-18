// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ZQLQuery 操作ZQLQuery
func (cli *ZSClient) ZQLQuery(params param.ZQLQueryParam) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.Get("v1/zql", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

