// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteMetricDataHttpReceiver deletes MetricDataHttpReceiver
func (cli *ZSClient) DeleteMetricDataHttpReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers/{uuid}", uuid, string(deleteMode))
}
