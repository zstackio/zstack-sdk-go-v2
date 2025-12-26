// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEmailAddressOfSNSEmailEndpoint deletes EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/emails/{endpointUuid}/email-addresses/{emailAddressUuid}", uuid, string(deleteMode))
}
