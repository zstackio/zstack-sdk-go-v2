// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// StartDataProtection starts DataProtection
func (cli *ZSClient) StartDataProtection(params param.StartDataProtectionParam) (*view.StartDataProtectionEventView, error) {
	resp := view.StartDataProtectionEventView{}
	if err := cli.Post("v1/start/data/protection/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtectionAsync Async
func (cli *ZSClient) StartDataProtectionAsync(params param.StartDataProtectionParam) (string, error) {

	resource := "v1/start/data/protection/"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
