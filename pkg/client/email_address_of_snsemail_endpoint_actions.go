// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEmailAddressOfSNSEmailEndpoint 删除EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/emails/{endpointUuid}/email-addresses/{emailAddressUuid}", uuid, string(deleteMode))
}

