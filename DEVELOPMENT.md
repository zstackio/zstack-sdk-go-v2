# ZStack SDK Go Development Standards

This document defines the development standards and coding conventions for the ZStack SDK Go project

---

## 1. Project Structure

```
zstack-sdk-go/
├── pkg/
│   ├── client/          # API client and action methods
│   ├── param/           # Request parameter structs
│   ├── view/            # Response view structs
│   ├── errors/          # Error definitions and handling
│   ├── util/            # Common utility packages
│   │   ├── jsonutils/   # JSON utilities
│   │   ├── httputils/   # HTTP utilities
│   │   └── ...          # Other utilities
│   └── test1/           # Integration tests
├── go.mod
├── go.sum
└── README.md
```

### Package Responsibilities

| Package | Responsibility | Naming Convention |
|---------|---------------|-------------------|
| `client` | API action method implementations | `{resource}_actions.go` |
| `param` | Request parameter definitions | `{resource}_params.go` |
| `view` | Response data structures | `{resource}_views.go` |
| `errors` | Error type definitions | `errors.go`, `consts.go` |
| `util` | Common utility functions | Organized by sub-packages |

---

## 2. Code Conventions

### 2.1 File Header

**All Go files must include a copyright notice:**

```go
// Copyright (c) ZStack.io, Inc.

package packagename
```

### 2.2 Import Order

Organize imports in the following order, separated by blank lines:

```go
import (
    // 1. Standard library
    "context"
    "fmt"
    "net/http"

    // 2. Third-party packages
    "github.com/kataras/golog"

    // 3. Internal project packages
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/errors"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)
```

---

## 3. Naming Conventions

### 3.1 File Naming

| Type | Format | Example |
|------|--------|---------|
| Actions | `{resource}_actions.go` | `vm_instance_actions.go` |
| Params | `{resource}_params.go` | `vm_instance_params.go` |
| Views | `{resource}_views.go` | `vm_instance_views.go` |
| Tests | `{resource}_test.go` | `vm_instance_test.go` |

### 3.2 Type Naming

```go
// Parameter structs: {Action}{Resource}Param
type CreateVmInstanceParam struct { ... }
type UpdateVmInstanceParam struct { ... }

// Detail parameters: {Action}{Resource}DetailParam
type CreateVmInstanceDetailParam struct { ... }

// View structs: {Resource}InventoryView or {Resource}View
type VmInstanceInventoryView struct { ... }
type VMConsoleAddressView struct { ... }

// Type aliases for enums
type DeleteMode string
type InstanceType string
```

### 3.3 Method Naming

```go
// CRUD operations
func (cli *ZSClient) Create{Resource}(params) (*View, error)
func (cli *ZSClient) Query{Resource}(params) ([]View, error)
func (cli *ZSClient) Get{Resource}(uuid) (*View, error)
func (cli *ZSClient) Update{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Destroy{Resource}(uuid, deleteMode) error
func (cli *ZSClient) Delete{Resource}(uuid, deleteMode) error

// Specific operations
func (cli *ZSClient) Start{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Stop{Resource}(uuid, params) (*View, error)
func (cli *ZSClient) Attach{A}To{B}(aUUID, bUUID) (*View, error)
func (cli *ZSClient) Detach{A}From{B}(aUUID, bUUID) (*View, error)
```

### 3.4 Constant Naming

```go
// Use type aliases for enums
type InstanceType string

const (
    UserVm      InstanceType = "UserVm"
    ApplianceVm InstanceType = "ApplianceVm"
)

// Error constants
const (
    ErrNotFound    = Error("NotFoundError")
    ErrDuplicateId = Error("DuplicateIdError")
)
```

---

## 4. Struct Design Patterns

### 4.1 Base Struct Embedding

**View structs use embedding to share common fields:**

```go
// Base info view
type BaseInfoView struct {
    UUID        string `json:"uuid"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

// Time info view
type BaseTimeView struct {
    CreateDate time.Time `json:"createDate"`
    LastOpDate time.Time `json:"lastOpDate"`
}

// Resource view embeds base structs
type VmInstanceInventoryView struct {
    BaseInfoView
    BaseTimeView
    
    ZoneUUID    string `json:"zoneUuid"`
    ClusterUUID string `json:"clusterUuid"`
    // ... other fields
}
```

### 4.2 Parameter Struct Embedding

```go
// Base parameters
type BaseParam struct {
    SystemTags []string `json:"systemTags,omitempty"`
    UserTags   []string `json:"userTags,omitempty"`
    RequestIp  string   `json:"requestIp,omitempty"`
}

// Request param embeds base param
type CreateVmInstanceParam struct {
    BaseParam
    Params CreateVmInstanceDetailParam `json:"params"`
}
```

### 4.3 Builder Pattern (Method Chaining)

```go
// Config builder
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

// Usage
client := client.NewZSClient(
    client.DefaultZSConfig("10.0.0.1").
        AccessKey("key-id", "key-secret").
        Debug(true),
)

// Query parameter builder
params := param.NewQueryParam().
    AddQ("name=test").
    Limit(10).
    Start(0)
```

---

## 5. JSON Tag Conventions

### 5.1 Field Tags

```go
type ExampleStruct struct {
    // Required fields: no omitempty
    UUID string `json:"uuid"`
    
    // Optional fields: use omitempty
    Description string `json:"description,omitempty"`
    
    // Pointer types to distinguish zero values from unset
    RootDiskSize *int64 `json:"rootDiskSize"`
    CpuNum       *int   `json:"cpuNum"`
}
```

### 5.2 Field Comments

**All exported fields must have descriptive comments:**

```go
type VmInstanceInventoryView struct {
    UUID             string `json:"uuid"`             // Resource UUID, unique identifier
    ZoneUUID         string `json:"zoneUuid"`         // Zone UUID
    ClusterUUID      string `json:"clusterUuid"`      // Cluster UUID
    MemorySize       int64  `json:"memorySize"`       // Memory size in bytes
    CPUNum           int    `json:"cpuNum"`           // Number of CPUs
}
```

---

## 6. Error Handling Conventions

### 6.1 Error Definition

```go
// Use custom error type
type Error string

func (e Error) Error() string {
    return string(e)
}

// Predefined error constants
const (
    ErrNotFound    = Error("NotFoundError")
    ErrDuplicateId = Error("DuplicateIdError")
    ErrParameter   = Error("ParameterError")
)
```

### 6.2 Error Wrapping

```go
import "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/errors"

// Use Wrap to add context
if err != nil {
    return errors.Wrap(err, "failed to create vm instance")
}

// Use Wrapf for formatted context
if err != nil {
    return errors.Wrapf(err, "failed to query %s", resource)
}
```

### 6.3 API Method Error Handling

```go
func (cli *ZSClient) GetVmInstance(uuid string) (*view.VmInstanceInventoryView, error) {
    var resp view.VmInstanceInventoryView
    if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
        return nil, err  // Return error directly, let caller handle
    }
    return &resp, nil
}
```

---

## 7. API Method Implementation Standards

### 7.1 Standard Method Template

```go
// {Description} Method description
func (cli *ZSClient) {MethodName}(params...) (*view.{ReturnType}, error) {
    var resp view.{ReturnType}
    if err := cli.{HttpMethod}("v1/{resource}", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}
```

### 7.2 Complete Examples

```go
// CreateVmInstance creates a VM instance
func (cli *ZSClient) CreateVmInstance(params param.CreateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
    resp := view.VmInstanceInventoryView{}
    if err := cli.Post("v1/vm-instances", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// QueryVmInstance queries VM instance list
func (cli *ZSClient) QueryVmInstance(params param.QueryParam) ([]view.VmInstanceInventoryView, error) {
    var resp []view.VmInstanceInventoryView
    return resp, cli.List("v1/vm-instances", &params, &resp)
}

// DestroyVmInstance deletes a VM instance
func (cli *ZSClient) DestroyVmInstance(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}
```

---

## 8. Testing Standards

### 8.1 Test File Location

Test files are placed in the `pkg/test1/` directory with naming format: `{resource}_test.go`

### 8.2 Test Function Naming

```go
func Test{MethodName}(t *testing.T) {
    // Test implementation
}
```

### 8.3 Test Template

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

## 9. Adding New Resource Support

When adding support for a new ZStack resource, follow these steps:

### Step 1: Define View Structs

Create `pkg/view/{resource}_views.go`:

```go
// Copyright (c) ZStack.io, Inc.

package view

type {Resource}InventoryView struct {
    BaseInfoView
    BaseTimeView
    
    // Resource-specific fields
    Field1 string `json:"field1"` // Field description
    Field2 int    `json:"field2"` // Field description
}
```

### Step 2: Define Parameter Structs

Create `pkg/param/{resource}_params.go`:

```go
// Copyright (c) ZStack.io, Inc.

package param

type Create{Resource}Param struct {
    BaseParam
    Params Create{Resource}DetailParam `json:"params"`
}

type Create{Resource}DetailParam struct {
    Name        string `json:"name"`        // Name
    Description string `json:"description"` // Description
    // Other parameters
}
```

### Step 3: Implement API Methods

Create `pkg/client/{resource}_actions.go`:

```go
// Copyright (c) ZStack.io, Inc.

package client

import (
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
    "github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// Create{Resource} creates a resource
func (cli *ZSClient) Create{Resource}(params param.Create{Resource}Param) (*view.{Resource}InventoryView, error) {
    resp := view.{Resource}InventoryView{}
    if err := cli.Post("v1/{resources}", params, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// Query{Resource} queries resource list
func (cli *ZSClient) Query{Resource}(params param.QueryParam) ([]view.{Resource}InventoryView, error) {
    var resp []view.{Resource}InventoryView
    return resp, cli.List("v1/{resources}", &params, &resp)
}

// Get{Resource} gets a single resource
func (cli *ZSClient) Get{Resource}(uuid string) (*view.{Resource}InventoryView, error) {
    var resp view.{Resource}InventoryView
    if err := cli.Get("v1/{resources}", uuid, nil, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}

// Destroy{Resource} deletes a resource
func (cli *ZSClient) Destroy{Resource}(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/{resources}", uuid, string(deleteMode))
}
```

### Step 4: Write Tests

Create integration tests in `pkg/test1/{resource}_test.go`.

---

## 10. Code Review Checklist

Before submitting code, ensure:

- [ ] All files include copyright notice
- [ ] Imports are organized in standard order
- [ ] All exported types and functions have comments
- [ ] Struct fields have JSON tags and comments
- [ ] Error handling correctly uses the errors package
- [ ] Naming follows project conventions
- [ ] Tests cover main functionality
- [ ] Code is formatted with `gofmt`

---

## 11. Go Version and Dependencies

- **Go Version**: 1.22.0+
- **Main Dependencies**:
  - `github.com/kataras/golog` - Logging
  - `github.com/pkg/errors` - Error handling
  - `github.com/fatih/color` - Terminal colors
  - `github.com/fatih/structs` - Struct reflection

---

*Document Version: 1.0*
*Last Updated: 2024-12*
