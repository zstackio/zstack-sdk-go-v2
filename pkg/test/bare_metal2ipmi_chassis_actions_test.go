// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateBareMetal2IpmiChassis(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2IpmiChassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2IpmiChassis found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2IpmiChassisParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2IpmiChassisParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2IpmiChassis(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2IpmiChassis error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2IpmiChassis result: %s", result.UUID)
}

func TestAddBareMetal2IpmiChassis(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddBareMetal2IpmiChassis requires valid creation parameters")

}
