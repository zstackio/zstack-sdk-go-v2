// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(uuid string, params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.UpdateEmailAddressOfSNSEmailEndpointEventView, error) {
	resp := view.UpdateEmailAddressOfSNSEmailEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/emails/email-addresses", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
