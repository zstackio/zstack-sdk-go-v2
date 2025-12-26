// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMetricDataHttpReceiver deletes MetricDataHttpReceiver
func (cli *ZSClient) DeleteMetricDataHttpReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers/{uuid}", uuid, string(deleteMode))
}
