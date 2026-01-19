// Copyright (c) ZStack.io, Inc.

package test

import (
	"fmt"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/util/jsonutils"
)

func TestQueryVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstance error: %v", err)
		return
	}
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s\t%d", r.UUID, r.Name, r.State, r.Platform, r.HypervisorType, r.MemorySize)
	}
	golog.Infof("======================================")
}

func TestQueryVmInstance2(t *testing.T) {
	// Query with conditions - similar to TestQueryImage2
	params := param.NewQueryParam()
	params.AddQ("state=Running")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryVmInstance(&params)
	if err != nil {
		t.Errorf("TestQueryVmInstance2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Running VMs:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%d CPU\t%d MB RAM", r.UUID, r.Name, r.State, r.HypervisorType, r.CpuNum, r.MemorySize/1024/1024)
	}
	golog.Infof("======================================")
}

func TestQueryVmInstance3(t *testing.T) {
	// Query with multiple conditions
	params := param.NewQueryParam()
	params.AddQ("state!=Destroyed")
	params.AddQ(fmt.Sprintf("platform=%s", "Linux"))
	params.Limit(5)
	result, err := accessKeyAuthCli.QueryVmInstance(&params)
	if err != nil {
		t.Errorf("TestQueryVmInstance3 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Linux VMs (not Destroyed):", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Platform)
	}
	golog.Infof("======================================")
}

func TestPageVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	queryParam.Start(0)
	result, total, err := accessKeyAuthCli.PageVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestPageVmInstance error: %v", err)
		return
	}
	golog.Infof("PageVmInstance result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%d", r.UUID, r.Name, r.State, r.Platform, total)
	}
}

func TestGetVmInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstance found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetVmInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmInstance error: %v", err)
		return
	}
	golog.Infof("GetVmInstance result: %s", result.UUID)
	golog.Info(jsonutils.Marshal(result))
}

func TestUpdateVmInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVmInstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accessKeyAuthCli.UpdateVmInstance(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVmInstance error: %v", err)
		return
	}
	golog.Infof("UpdateVmInstance result: %s", result.UUID)
}

/*
func TestCreateVmInstance(t *testing.T) {
	// WARNING: This test will create a real resource!
	//	t.Skip("TestCreateVmInstance is skipped by default. Implement with valid params to test creation.")

	network, err := accessKeyAuthCli.QueryL3Network(&param.QueryParam{})
	if err != nil {
		golog.Errorf("TestZSClient_QueryL3Network %v: %v", network, err)
	}
	images, err := accessKeyAuthCli.QueryImage(&param.QueryParam{})
	if err != nil {
		golog.Errorf("TestZSClient_QueryImage %v: %v", images, err)
	}
	rootDiskSize := int64(1073741824)
	memSize := int64(1073741824)
	cpuNum := 1
	instance, err := accessKeyAuthCli.CreateVmInstance(param.CreateVmInstanceParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"resourceConfig::vm::vm.clock.track::guest", "cdroms::Empty::None::None"},
			UserTags:   nil,
			RequestIp:  "",
		},
		Params: param.CreateVmInstanceParamDetail{
			Name: "test-uuid",
			//InstanceOfferingUUID:            "",
			ImageUuid:            &images[0].UUID,
			L3NetworkUuids:       []string{network[0].UUID},
			RootDiskSize:         &rootDiskSize,
			DataDiskSizes:                   []int64{10240},
			DefaultL3NetworkUuid:            &network[0].UUID,
			TagUuids:   nil,
			MemorySize: &memSize,
			CpuNum:     &cpuNum,
		},
	})
	if err != nil {
		t.Errorf("TestCreateVmInstance %v", err)
	}
	golog.Println(instance)

}

func TestCloneVmInstance(t *testing.T) {
	// Clone operation
	// t.Skip("TestCloneVmInstance requires a valid resource to clone")
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Stopped")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestCloneVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Stopped VmInstance found to test Clone")
		return
	}

	cloneParam := param.CloneVmInstanceParam{
		Params: param.CloneVmInstanceParamDetail{
			Names:    []string{"cloned-vm-" + list[0].UUID[:8]},
		},
	}
	result, err := accessKeyAuthCli.CloneVmInstance(list[0].UUID, cloneParam)
	if err != nil {
		t.Errorf("TestCloneVmInstance error: %v", err)
		return
	}
	golog.Infof("CloneVmInstance result: %s", result.UUID)
}
*/

func TestResumeVmInstance(t *testing.T) {
	// ResumeVmInstance operation
	// t.Skip("TestResumeVmInstance requires manual implementation")
	// Workflow: Running -> Pause -> Resume
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Running")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestResumeVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Running VmInstance found to test Pause/Resume")
		return
	}
	// 1. Pause
	pauseParam := param.PauseVmInstanceParam{}
	_, err = accessKeyAuthCli.PauseVmInstance(list[0].UUID, pauseParam)
	if err != nil {
		t.Errorf("TestResumeVmInstance Pause error: %v", err)
		return
	}

	// 2. Resume
	resumeParam := param.ResumeVmInstanceParam{}
	result, err := accessKeyAuthCli.ResumeVmInstance(list[0].UUID, resumeParam)
	if err != nil {
		t.Errorf("TestResumeVmInstance Resume error: %v", err)
		return
	}
	golog.Infof("ResumeVmInstance result: %s", result.UUID)
}

func TestStartVmInstance(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	//t.Skip("TestStartVmInstance requires a stopped resource UUID")
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Stopped")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstance Query error: %v", err)
		return
	}

	if len(list) == 0 {
		t.Skip("No VmInstance found to test Start")
		return
	}

	resp, err := accessKeyAuthCli.StartVmInstance(list[0].UUID,
		param.StartVmInstanceParam{
			Params: param.StartVmInstanceParamDetail{
				//	HostUuid: "a70a21e621394573bf10bd3749045bc9",
			},
		})
	if err != nil {
		t.Errorf("TestStartVmInstance : %v", err)
	}
	golog.Println(resp)
}

func TestStopVmInstance(t *testing.T) {
	// Stop operation - requires a running resource UUID
	//t.Skip("TestStopVmInstance requires a running resource UUID")
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Running")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil || len(list) == 0 {
		t.Skip("No running  found")
		return
	}
	stopParam := param.StopVmInstanceParam{}
	result, err := accessKeyAuthCli.StopVmInstance(list[0].UUID, stopParam)
	if err != nil {
		t.Errorf("TestStopVmInstance error: %v", err)
	}
	golog.Infof("StopVmInstance result: %v", result)

}

func TestExpungeVmInstance(t *testing.T) {
	// Expunge operation - permanently deletes
	// t.Skip("TestExpungeVmInstance is dangerous - permanently deletes resource")
	// Workflow: Destroy -> Expunge

	// 1. Find a non-destroyed VM (Running or Stopped)
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state!=Destroyed")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestExpungeVmInstance query error: %v", err)
		return
	}
	if len(list) == 0 {
		// Try finding a Destroyed one if no active ones exist
		queryParam = param.NewQueryParam()
		queryParam.AddQ("state=Destroyed")
		queryParam.Limit(1)
		list, err = accessKeyAuthCli.QueryVmInstance(&queryParam)
		if len(list) == 0 {
			t.Skip("No VmInstance found to test Expunge")
			return
		}
	} else {
		// Destroy it first
		err = accessKeyAuthCli.DestroyVmInstance(list[0].UUID, param.DeleteModePermissive)
		if err != nil {
			t.Errorf("TestExpungeVmInstance Destroy error: %v", err)
			return
		}
	}

	// 2. Expunge
	err = accessKeyAuthCli.ExpungeVmInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestExpungeVmInstance error: %v", err)
		return
	}
	golog.Infof("ExpungeVmInstance result: %s", list[0].UUID)
}

func TestRebootVmInstance(t *testing.T) {
	// Reboot operation - requires a running resource
	// t.Skip("TestRebootVmInstance requires a running resource UUID")
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Running")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestRebootVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Running VmInstance found to test Reboot")
		return
	}

	rebootParam := param.RebootVmInstanceParam{}
	result, err := accessKeyAuthCli.RebootVmInstance(list[0].UUID, rebootParam)
	if err != nil {
		t.Errorf("TestRebootVmInstance error: %v", err)
		return
	}
	golog.Infof("RebootVmInstance result: %s", result.UUID)
}

func TestDestroyVmInstance(t *testing.T) {
	// DestroyVmInstance operation
	// t.Skip("TestDestroyVmInstance requires manual implementation")
	queryParam := param.NewQueryParam()
	// Can destroy Running or Stopped VMs
	queryParam.AddQ("state!=Destroyed")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestDestroyVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No valid VmInstance found to test Destroy")
		return
	}

	err = accessKeyAuthCli.DestroyVmInstance(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDestroyVmInstance error: %v", err)
		return
	}
	golog.Infof("DestroyVmInstance result: %s", list[0].UUID)
}

func TestRecoverVmInstance(t *testing.T) {
	// Recover operation - requires a deleted resource
	// t.Skip("TestRecoverVmInstance requires a deleted resource UUID")
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Destroyed")
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestRecoverVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Destroyed VmInstance found to test Recover")
		return
	}

	recoverParam := param.RecoverVmInstanceParam{}
	result, err := accessKeyAuthCli.RecoverVmInstance(list[0].UUID, recoverParam)
	if err != nil {
		t.Errorf("TestRecoverVmInstance error: %v", err)
		return
	}
	golog.Infof("RecoverVmInstance result: %s", result.UUID)
}
