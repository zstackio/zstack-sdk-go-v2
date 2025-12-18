// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanV2VConversionCache operates on CleanV2VConversionCache
func (cli *ZSClient) CleanV2VConversionCache(uuid string, params param.CleanV2VConversionCacheParam) (*view.CleanV2VConversionCacheEventView, error) {
	resp := view.CleanV2VConversionCacheEventView{}
	if err := cli.Put("v1/v2v/conversion/host/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
