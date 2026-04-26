# ZStack SDK for Go



## ZStack SDK for Go

This repository provides a Go SDK for interacting with the ZStack cloud platform API. It allows developers to manage ZStack resources programmatically using Go.

### Features

- Complete API coverage for ZStack cloud platform
- Type-safe Go client for ZStack API operations
- Support for authentication and session management
- Structured request and response handling
- Error handling and logging capabilities
- Compatible with Terraform provider development

### Installation

```bash
go get github.com/zstackio/zstack-sdk-go-v2
```

### Usage

#### Initialize Client with Account Login

```go
import (
    "context"

    "github.com/zstackio/zstack-sdk-go-v2/pkg/client"
    "github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

// Create a client with account login authentication
accountLoginClient := client.NewZSClient(
    client.DefaultZSConfig("ZStack MN HOST IP").
        LoginAccount("admin", "password").
        ReadOnly(false).
        Debug(true),
)

// Login to ZStack
_, err := accountLoginClient.Login(context.Background())
if err != nil {
    // Handle error
}
defer accountLoginClient.Logout(context.Background())

// Query clusters
clusters, err := accountLoginClient.QueryCluster(param.NewQueryParam())
if err != nil {
    // Handle error
}
// Use clusters data

```

#### Initialize Client with Access Key
```go
import (
    "github.com/zstackio/zstack-sdk-go-v2/pkg/client"
    "github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

// Create a client with access key authentication
accessKeyClient := client.NewZSClient(
    client.DefaultZSConfig("ZStack MN HOST IP").
        AccessKey("your-access-key-id", "your-access-key-secret").
        ReadOnly(false).
        Debug(false),
)

// Query clusters
clusters, err := accessKeyClient.QueryCluster(param.NewQueryParam())
if err != nil {
    // Handle error
}
// Use clusters data

```
### Authentication

The SDK supports the following authentication methods:
- Username and password authentication
- Session-based authentication with automatic token renewal
- API key authentication

### API Coverage

This SDK provides Go bindings for all major ZStack API operations, including but not limited to:

- Compute resources (instances, volumes)
- Network resources (L2/L3 networks, security groups)
- Storage resources (primary storage, backup storage)
- Identity and access management
- System management operations


### Testing

The repository now distinguishes between fast default checks and environment-backed test suites.

#### Default verification

Run the full repository checks without requiring a live ZStack environment:

```bash
go test ./...
```

#### Environment-backed test suites

The following test packages are disabled by default and only run when their gate variable is set:

| Package | Enable variable | Purpose |
| --- | --- | --- |
| `pkg/test` | `ZSTACK_ENABLE_PKG_TEST=1` | Auto-generated SDK action tests |
| `pkg/integration_test` | `ZSTACK_ENABLE_INTEGRATION_TEST=1` | Integration tests against a live API endpoint |
| `pkg/real_test` | `ZSTACK_ENABLE_REAL_TEST=1` | Real-environment regression coverage |

Examples:

```bash
ZSTACK_ENABLE_PKG_TEST=1 go test ./pkg/test
ZSTACK_ENABLE_INTEGRATION_TEST=1 ZSTACK_HOST=127.0.0.1 ZSTACK_PORT=8080 ZSTACK_ACCOUNT=admin ZSTACK_PASSWORD=password go test ./pkg/integration_test
ZSTACK_ENABLE_REAL_TEST=1 go test ./pkg/real_test
```

To run tests and generate an HTML report:

**Windows (PowerShell):**
```powershell
./scripts/run_tests_with_report.ps1
```

**Linux (Bash):**
```bash
./scripts/run_tests_with_report.sh
```

### Contributing


Contributions to the ZStack Go SDK are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

### Related Projects

- [Terraform Provider for ZStack](https://github.com/terraform-zstack-modules/terraform-provider-zstack)
- [ZStack Documentation](https://www.zstack.io/help/dev_manual/dev_guide/v5/)

### Support

For issues, questions and discussions please use the [GitHub Issues](https://github.com/zstackio/zstack-sdk-go-v2/issues).
