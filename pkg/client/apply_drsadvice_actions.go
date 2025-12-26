// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ApplyDRSAdvice operates on ApplyDRSAdvice
func (cli *ZSClient) ApplyDRSAdvice(uuid string, params param.ApplyDRSAdviceParam) (*view.ApplyDRSAdviceEventView, error) {
	resp := view.ApplyDRSAdviceEventView{}
	if err := cli.Put("v1/clusters/drs/advice/{adviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
