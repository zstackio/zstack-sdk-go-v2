// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(uuid string, params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.UpdateEmailAddressOfSNSEmailEndpointEventView, error) {
	resp := view.UpdateEmailAddressOfSNSEmailEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/emails/email-addresses", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
