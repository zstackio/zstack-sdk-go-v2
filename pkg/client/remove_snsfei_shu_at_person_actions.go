// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSNSFeiShuAtPerson removes SNSFeiShuAtPerson
func (cli *ZSClient) RemoveSNSFeiShuAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/feishu/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}
