// Copyright (c) ZStack.io, Inc.

package test

import (
    "testing"
    "github.com/kataras/golog"
    "github.com/zstackio/zstack-sdk-go-v2/pkg/util/jsonutils"
)

// TestGetVersion 测试获取ZStack版本信息
// 这个API没有参数，调用时不需要传入任何参数
func TestGetVersion(t *testing.T) {
    golog.Infof("Calling GetVersion...")

    result, err := accessKeyAuthCli.GetVersion()
    if err != nil {
        t.Errorf("TestGetVersion error: %v", err)
        return
    }

    golog.Infof("======================================")
    golog.Infof("ZStack Version: %s", result.Version)
    golog.Info(jsonutils.Marshal(result))
    golog.Infof("======================================")
}

// TestGetLicenseInfo 测试获取许可证信息
// 这个API也没有参数，直接调用即可
func TestGetLicenseInfo(t *testing.T) {
    golog.Infof("Calling GetLicenseInfo...")
    
    result, err := accessKeyAuthCli.GetLicenseInfo()
    if err != nil {
        t.Errorf("TestGetLicenseInfo error: %v", err)
        return
    }

    golog.Infof("======================================")
    golog.Infof("License Info:")
    golog.Infof("UUID: %s", result.UUID)
    golog.Info(jsonutils.Marshal(result))
    golog.Infof("======================================")
}
