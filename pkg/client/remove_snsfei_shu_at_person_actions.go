// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSNSFeiShuAtPerson removes SNSFeiShuAtPerson
func (cli *ZSClient) RemoveSNSFeiShuAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/feishu/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}
