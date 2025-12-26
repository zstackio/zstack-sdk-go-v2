// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSNSDingTalkAtPerson removes SNSDingTalkAtPerson
func (cli *ZSClient) RemoveSNSDingTalkAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/ding-talk/{endpointUuid}/at-persons/{phoneNumber}", uuid, string(deleteMode))
}
