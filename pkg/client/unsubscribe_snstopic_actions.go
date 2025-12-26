// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// UnsubscribeSNSTopic operates on UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", uuid, string(deleteMode))
}
