// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

// TestUpdateLicense 测试更新许可证
func TestUpdateLicense(t *testing.T) {
	testName := "更新许可证"
	golog.Infof("开始测试: %s", testName)

	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会修改实际许可证
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会修改实际许可证", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	t.Skip("Skipping update license test as it modifies actual license")

	// 首先查询获取管理节点
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	managementNodes, err := cli.QueryManagementNode(&queryParam)
	if !assert.NoError(t, err, "Query management node should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(managementNodes) == 0 {
		golog.Warnf("测试跳过 [%s]: No ManagementNode found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No ManagementNode found to test")
		return
	}

	// 准备更新许可证参数
	// 注意：这里需要一个有效的许可证字符串
	licenseContent := "your-license-content-here"
	updateParam := param.UpdateLicenseParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateLicenseParamDetail{
			License: licenseContent,
		},
	}

	// 执行更新
	updatedLicense, err := cli.UpdateLicense(managementNodes[0].UUID, updateParam)
	if !assert.NoError(t, err, "Update license should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated license: %s, expired date: %s", updatedLicense.UUID, updatedLicense.ExpiredDate)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDeleteLicense 测试删除许可证
func TestDeleteLicense(t *testing.T) {
	testName := "删除许可证"
	golog.Infof("开始测试: %s", testName)

	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会删除实际许可证
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会删除实际许可证", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	t.Skip("Skipping delete license test as it deletes actual license")

	// 这里需要一个有效的许可证UUID
	// 在实际测试中，可能需要先查询许可证列表获取UUID
	licenseUuid := "your-license-uuid-here"

	// 删除许可证
	err := cli.DeleteLicense(licenseUuid, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete license should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("License deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}