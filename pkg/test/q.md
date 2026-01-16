# SDK 问题清单及状态

> 测试时间: 2026-01-16 17:06  
> 测试环境: 172.24.249.239

| ID | 模块/功能 | 状态 | 说明/代码示例 | 测试结果 |
|:---:|:---|:---:|:---|:---|
| 1 | `GetAccountResourceRef` | ⚠️ 待确认 | API 是否支持 Get 方法需确认。<br>代码逻辑已实现：`cli.Get("v1/accounts/resources/refs", ...)` | 待确认 ZStack API 支持情况 |
| 2 | `GetGlobalConfig` | ❌ 需修复 | API 路径错误。<br>当前：`v1/global-configurations/{uuid}`<br>应为：`v1/resource-configurations/{resourceUuid}/{category}/{name}` | `DuplicateIdError` |
| 3.1 | `PageXXX` 函数 | ✅ 已实现 | `PageImage` 已增加。<br>实现逻辑：调用 `cli.Page("v1/images", ...)` | TestPageImage ✅ PASS |
| 3.2 | Post/Put 自动解包 | ✅ 正常工作 | Client 自动处理 `inventory` 字段，`AddImage` 无需解包。<br>`AddImageParam` 需使用 `json:"params"` | TestAddImage ✅ PASS |
| 3.3 | `ExpungeXXX` | ✅ 已实现 | 统一修改为 Put 请求。<br>示例：`cli.Put("v1/images", uuid, ...)` | 逻辑已修正 |
| 4 | View 结构体 | ✅ 已实现 | 移除指针，改用值类型。<br>示例：`Description string` 代替 `*string` | Build 正常 |
| 5 | `UpdateXXXParam` JSON Tag | ✅ 已修复 | Tag 修正为功能名称。<br>示例：`json:"updateImage"` | TestUpdateImage ✅ PASS |
| 6 | `ChangeXXXStateParam` JSON Tag | ✅ 已修复 | Tag 修正为功能名称。<br>示例：`json:"changeImageState"` | TestChangeImageState ✅ PASS |
| 7 | `ChangeXXXState` UUID 参数 | ⚠️ 冗余 | 函数签名中 `uuid` 与 `Params.Uuid` 重复。<br>示例：`ChangeImageState(uuid, param)` | 测试通过，但参数冗余 |
| 8 | 其他参数 (Clone/Start/Set/Change) | ⚠️ 待检查 | 类似于 Update/Change 的 JSON Tag 问题需全面检查。<br>如 `CloneVmInstance` 等 | 已增加测试覆盖，全部验证通过 |
| 9 | `CreateVmInstance` | ✅ 已验证 | 创建 VM API 调用成功。<br>**关键修正**: 参数 Struct 必须使用 `json:"params"` Tag。<br>需注意设置 go test 超时 (如 `-timeout 30m`) | TestCreateVmInstance ✅ PASS<br>(耗时约13s) |
| 10 | `Stop`/`Start`/`Reboot`... | ✅ 已验证 | 大部分 VM 动作函数测试通过。 | 见下方详细测试结果 |
| 11 | `RecoverVmInstance` | ❌ 需修复 | 请求体发送 `{}` 而非 `{"recoverVmInstance":{}}`<br>服务端返回: `400 - body doesn't contain action mapping` | TestRecoverVmInstance ❌ FAIL |

## 详细测试结果

| 测试项 | 状态 | 备注 |
|:---|:---:|:---|
| TestQueryImage1 | ✅ PASS | |
| TestGetImage | ✅ PASS | |
| TestUpdateImage | ✅ PASS | |
| TestPageImage | ✅ PASS | |
| TestChangeImageState | ✅ PASS | |
| TestAddImage | ✅ PASS | JSON Tag 要修改为 `params` |
| TestGetGlobalConfig | ❌ FAIL | 报错 `DuplicateIdError` |
| TestQueryVmInstance | ✅ PASS | |
| TestGetVmInstance | ✅ PASS | |
| TestUpdateVmInstance | ✅ PASS | |
| TestCreateVmInstance | ✅ PASS | JSON Tag 要修改为 `params`, 超时需调整 |
| TestCloneVmInstance | ✅ PASS | 创建并启动了新 VM |
| TestRebootVmInstance | ✅ PASS | (依赖 Running VM) |
| TestStartVmInstance | ✅ PASS | 启动了 Stopped 的 VM |
| TestStopVmInstance | ✅ PASS | 成功停止 VM |
| TestDestroyVmInstance | ✅ PASS | 成功删除 VM |
| TestExpungeVmInstance | ✅ PASS | 验证流程: Destroy -> Expunge |
| TestResumeVmInstance | ⏭️ SKIP | 无 Paused 状态 VM 可测试 |
| TestRecoverVmInstance | ❌ FAIL | 请求体错误: 发送 `{}` 而非 `{"recoverVmInstance":{}}` |

## 关键参数应要修正的记录

### 1. AddImageParam
```go
// AddImageParam AddImage request param
type AddImageParam struct {
	BaseParam
	Params AddImageParamDetail `json:"params"` // 原为 json:"addImage" -> 错误
}
```

### 2. CreateVmInstanceParam
```go
type CreateVmInstanceParam struct {
	BaseParam
	Params CreateVmInstanceParamDetail `json:"params"` // 原为 json:"createVmInstance" -> 错误
}
```

### 3. StartVmInstance 函数需要构造put，例子如下
```go
// StartVmInstance starts VmInstance
func (cli *ZSClient) StartVmInstance(uuid string, params param.StartVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Put("v1/vm-instances", uuid, map[string]struct{}{
		"startVmInstance": {}}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
```

### 4. StartVmInstanceParamDetail 应该要删除Uuid， 和StartVmInstance函数的uuid参数一样, 重复了
```go
// StartVmInstanceParamDetail StartVmInstance detail param
type StartVmInstanceParamDetail struct {
	//Uuid        string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"` // 仅保留其他参数
	HostUuid    string `json:"hostUuid,omitempty"`
}
```

### 5. Pause/Resume/Recover 参数需填充 Uuid
尽管函数签名中有 `uuid` 参数，但对应的 `Params` 结构体内若包含 `Uuid` 字段且为 required，则必须在 Params 中再次填充该 Uuid，否则报错 400 (Body doesn't contain action mapping)。
```go
pauseParam := param.PauseVmInstanceParam{
    Params: param.PauseVmInstanceParamDetail{ Uuid: uuid },
}
```
这也适用于 `RecoverVmInstanceParam` 和 `ResumeVmInstanceParam`。
