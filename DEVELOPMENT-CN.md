# ZStack SDK Go 开发规范与标准

ZStack SDK Go 项目的开发规范和编码标准。

---

## 1. 项目结构

```
zstack-sdk-go/
├── pkg/
│   ├── client/          # API 客户端和操作方法
│   ├── param/           # 请求参数结构体
│   ├── view/            # 响应视图结构体
│   ├── errors/          # 错误定义和处理
│   ├── util/            # 通用工具包
│   │   ├── jsonutils/   # JSON 处理
│   │   ├── httputils/   # HTTP 工具
│   │   └── ...          # 其他工具
│   └── test1/           # 集成测试
├── go.mod
├── go.sum
└── README.md
```

### 包职责说明

| 包名 | 职责 | 命名规范 |
|------|------|----------|
| `client` | API 操作方法实现 | `{resource}_actions.go` |
| `param` | 请求参数定义 | `{resource}_params.go` |
| `view` | 响应数据结构 | `{resource}_views.go` |
| `errors` | 错误类型定义 | `errors.go`, `consts.go` |
| `util` | 通用工具函数 | 按功能划分子包 |

---

## 2. 代码规范

### 2.1 文件头部

**所有 Go 文件必须包含版权声明：**

```go
// Copyright (c) ZStack.io, Inc.

package packagename
```

### 2.2 导入顺序

按以下顺序组织导入，组间用空行分隔：

```go
import (
    // 1. 标准库
    "context"
    "fmt"
    "net/http"

    // 2. 第三方库
    "github.com/kataras/golog"

    // 3. 项目内部包
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/errors"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)
```

---

## 3. 命名规范

### 3.1 文件命名

| 类型 | 格式 | 示例 |
|------|------|------|
| Actions | `{resource}_actions.go` | `vm_instance_actions.go` |
| Params | `{resource}_params.go` | `vm_instance_params.go` |
| Views | `{resource}_views.go` | `vm_instance_views.go` |
| Tests | `{resource}_test.go` | `vm_instance_test.go` |

### 3.2 类型命名

```go
// 参数结构体：{Action}{Resource}Param
type CreateVmInstanceParam struct { ... }
type UpdateVmInstanceParam struct { ... }

// 详细参数：{Action}{Resource}DetailParam
type CreateVmInstanceDetailParam struct { ... }

// 视图结构体：{Resource}InventoryView 或 {Resource}View
type VmInstanceInventoryView struct { ... }
type VMConsoleAddressView struct { ... }

// 类型别名
type DeleteMode string
type InstanceType string
```

### 3.3 方法命名

```go
// CRUD 操作
func (cli *ZSClient) Create{Resource}(params) (*View, error)
func (cli *ZSClient) Query{Resource}(params) ([]View, error)
func (cli *ZSClient) Get{Resource}(uuid) (*View, error)
func (cli *ZSClient) Update{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Destroy{Resource}(uuid, deleteMode) error
func (cli *ZSClient) Delete{Resource}(uuid, deleteMode) error

// 特定操作
func (cli *ZSClient) Start{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Stop{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Attach{A}To{B}(aUUID, bUUID) (*View, error)
func (cli *ZSClient) Detach{A}From{B}(aUUID, bUUID) (*View, error)
```

### 3.4 常量命名

```go
// 使用类型别名定义枚举
type InstanceType string

const (
    UserVm      InstanceType = "UserVm"
    ApplianceVm InstanceType = "ApplianceVm"
)

// 错误常量
const (
    ErrNotFound    = Error("NotFoundError")
    ErrDuplicateId = Error("DuplicateIdError")
)
```

---

## 4. 结构体设计模式

### 4.1 基础结构体嵌入

**View 结构体使用嵌入共享通用字段：**

```go
// 基础信息视图
type BaseInfoView struct {
    UUID        string `json:"uuid"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

// 时间信息视图
type BaseTimeView struct {
    CreateDate time.Time `json:"createDate"`
    LastOpDate time.Time `json:"lastOpDate"`
}

// 资源视图嵌入基础结构体
type VmInstanceInventoryView struct {
    BaseInfoView
    BaseTimeView
    
    ZoneUUID    string `json:"zoneUuid"`
    ClusterUUID string `json:"clusterUuid"`
    // ... 其他字段
}
```

### 4.2 参数结构体嵌入

```go
// 基础参数
type BaseParam struct {
    SystemTags []string `json:"systemTags,omitempty"`
    UserTags   []string `json:"userTags,omitempty"`
    RequestIp  string   `json:"requestIp,omitempty"`
}

// 请求参数嵌入基础参数
type CreateVmInstanceParam struct {
    BaseParam
    Params CreateVmInstanceDetailParam `json:"params"`
}
```

### 4.3 Builder 模式 (方法链)

```go
// 配置构建器
func DefaultZSConfig(hostname string) *ZSConfig {
    return NewZSConfig(hostname, defaultZStackPort, defaultZStackContextPath)
}

func (config *ZSConfig) AccessKey(id, secret string) *ZSConfig {
    config.accessKeyId = id
    config.accessKeySecret = secret
    config.authType = AuthTypeAccessKey
    return config
}

func (config *ZSConfig) Debug(debug bool) *ZSConfig {
    config.debug = debug
    return config
}

// 使用方式
client := client.NewZSClient(
    client.DefaultZSConfig("10.0.0.1").
        AccessKey("key-id", "key-secret").
        Debug(true),
)

// 查询参数构建器
params := param.NewQueryParam().
    AddQ("name=test").
    Limit(10).
    Start(0)
```

---

## 5. JSON 标签规范

### 5.1 字段标签

```go
type ExampleStruct struct {
    // 必填字段：无 omitempty
    UUID string `json:"uuid"`
    
    // 可选字段：使用 omitempty
    Description string `json:"description,omitempty"`
    
    // 指针类型用于区分零值和未设置
    RootDiskSize *int64 `json:"rootDiskSize"`
    CpuNum       *int   `json:"cpuNum"`
}
```

### 5.2 字段注释

**所有导出字段必须有中文注释说明：**

```go
type VmInstanceInventoryView struct {
    UUID             string `json:"uuid"`             // 资源UUID，唯一标识
    ZoneUUID         string `json:"zoneUuid"`         // 区域UUID
    ClusterUUID      string `json:"clusterUuid"`      // 集群UUID
    MemorySize       int64  `json:"memorySize"`       // 内存大小（字节）
    CPUNum           int    `json:"cpuNum"`           // CPU数量
}
```

---

## 6. 错误处理规范

### 6.1 错误定义

```go
// 使用自定义错误类型
type Error string

func (e Error) Error() string {
    return string(e)
}

// 预定义错误常量
const (
    ErrNotFound    = Error("NotFoundError")
    ErrDuplicateId = Error("DuplicateIdError")
    ErrParameter   = Error("ParameterError")
)
```

### 6.2 错误包装

```go
import "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/errors"

// 使用 Wrap 添加上下文
if err != nil {
    return errors.Wrap(err, "failed to create vm instance")
}

// 使用 Wrapf 格式化上下文
if err != nil {
    return errors.Wrapf(err, "failed to query %s", resource)
}
```

### 6.3 API 方法错误处理

```go
func (cli *ZSClient) GetVmInstance(uuid string) (*view.VmInstanceInventoryView, error) {
    var resp view.VmInstanceInventoryView
    if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
        return nil, err  // 直接返回错误，由调用方处理
    }
    return &resp, nil
}
```

---

## 7. API 方法实现规范

### 7.1 标准方法模板

```go
// {Description} 方法描述
func (cli *ZSClient) {MethodName}(params...) (*view.{ReturnType}, error) {
    var resp view.{ReturnType}
    if err := cli.{HttpMethod}("v1/{resource}", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}
```

### 7.2 完整示例

```go
// CreateVmInstance 创建虚拟机实例
func (cli *ZSClient) CreateVmInstance(params param.CreateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
    resp := view.VmInstanceInventoryView{}
    if err := cli.Post("v1/vm-instances", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// QueryVmInstance 查询虚拟机实例列表
func (cli *ZSClient) QueryVmInstance(params param.QueryParam) ([]view.VmInstanceInventoryView, error) {
    var resp []view.VmInstanceInventoryView
    return resp, cli.List("v1/vm-instances", &params, &resp)
}

// DestroyVmInstance 删除虚拟机实例
func (cli *ZSClient) DestroyVmInstance(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}
```

---

## 8. 测试规范

### 8.1 测试文件位置

测试文件放在 `pkg/test1/` 目录下，命名格式：`{resource}_test.go`

### 8.2 测试函数命名

```go
func Test{MethodName}(t *testing.T) {
    // 测试实现
}
```

### 8.3 测试模板

```go
// Copyright (c) ZStack.io, Inc.

package test

import (
    "testing"

    "github.com/kataras/golog"

    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/util/jsonutils"
)

func TestQueryVmInstance(t *testing.T) {
    data, err := accessKeyAuthCli.QueryVmInstance(param.NewQueryParam())
    if err != nil {
        t.Errorf("TestQueryVmInstance: %v", err)
    }
    golog.Info(jsonutils.Marshal(data))
}

func TestGetVmInstance(t *testing.T) {
    data, err := accountLoginCli.GetVmInstance("uuid-here")
    if err != nil {
        t.Errorf("TestGetVmInstance: %v", err)
    }
    golog.Info(jsonutils.Marshal(data))
}
```

---

## 9. 新增资源开发流程

当需要添加新的 ZStack 资源支持时，按以下步骤进行：

### 步骤 1：定义视图结构体

在 `pkg/view/{resource}_views.go` 中定义：

```go
// Copyright (c) ZStack.io, Inc.

package view

type {Resource}InventoryView struct {
    BaseInfoView
    BaseTimeView
    
    // 资源特定字段
    Field1 string `json:"field1"` // 字段说明
    Field2 int    `json:"field2"` // 字段说明
}
```

### 步骤 2：定义参数结构体

在 `pkg/param/{resource}_params.go` 中定义：

```go
// Copyright (c) ZStack.io, Inc.

package param

type Create{Resource}Param struct {
    BaseParam
    Params Create{Resource}DetailParam `json:"params"`
}

type Create{Resource}DetailParam struct {
    Name        string `json:"name"`        // 名称
    Description string `json:"description"` // 描述
    // 其他参数
}
```

### 步骤 3：实现 API 方法

在 `pkg/client/{resource}_actions.go` 中实现：

```go
// Copyright (c) ZStack.io, Inc.

package client

import (
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// Create{Resource} 创建资源
func (cli *ZSClient) Create{Resource}(params param.Create{Resource}Param) (*view.{Resource}InventoryView, error) {
    resp := view.{Resource}InventoryView{}
    if err := cli.Post("v1/{resources}", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// Query{Resource} 查询资源列表
func (cli *ZSClient) Query{Resource}(params param.QueryParam) ([]view.{Resource}InventoryView, error) {
    var resp []view.{Resource}InventoryView
    return resp, cli.List("v1/{resources}", &params, &resp)
}

// Get{Resource} 获取单个资源
func (cli *ZSClient) Get{Resource}(uuid string) (*view.{Resource}InventoryView, error) {
    var resp view.{Resource}InventoryView
    if err := cli.Get("v1/{resources}", uuid, nil, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// Destroy{Resource} 删除资源
func (cli *ZSClient) Destroy{Resource}(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/{resources}", uuid, string(deleteMode))
}
```

### 步骤 4：编写测试

在 `pkg/test1/{resource}_test.go` 中编写集成测试。

---

## 10. 代码审查清单

在提交代码前，请确保：

- [ ] 所有文件包含版权声明
- [ ] 导入按标准顺序组织
- [ ] 所有导出的类型和函数有注释
- [ ] 结构体字段有 JSON 标签和注释
- [ ] 错误处理正确使用 errors 包
- [ ] 命名遵循项目规范
- [ ] 测试覆盖主要功能
- [ ] 代码格式化通过 `gofmt`

---

## 11. Go 版本和依赖

- **Go 版本**: 1.22.0+
- **主要依赖**:
  - `github.com/kataras/golog` - 日志
  - `github.com/pkg/errors` - 错误处理
  - `github.com/fatih/color` - 终端颜色
  - `github.com/fatih/structs` - 结构体反射

---

*文档版本: 1.0*
*最后更新: 2024-12*
