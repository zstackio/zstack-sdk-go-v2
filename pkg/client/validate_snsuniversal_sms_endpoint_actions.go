// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateSNSUniversalSmsEndpoint operates on ValidateSNSUniversalSmsEndpoint
func (cli *ZSClient) ValidateSNSUniversalSmsEndpoint(uuid string, params param.ValidateSNSUniversalSmsEndpointParam) (*view.ValidateSNSApplicationEndpointEventView, error) {
	resp := view.ValidateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/validate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
