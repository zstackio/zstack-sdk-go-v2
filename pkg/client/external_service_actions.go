// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// ReloadExternalService operates on ReloadExternalService
func (cli *ZSClient) ReloadExternalService(params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.PutWithRespKey("v1/external/services", "", "", map[string]interface{}{
		"reloadExternalService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExternalServices gets ExternalServices by uuid
func (cli *ZSClient) GetExternalServices() (*view.ExternalServiceInventoryView, error) {
	var resp view.ExternalServiceInventoryView
	if err := cli.GetWithRespKey("v1/external/services", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
