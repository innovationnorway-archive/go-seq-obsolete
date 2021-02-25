# \LicensesApi

All URIs are relative to *https://localhost:80*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DowngradeLicense**](LicensesApi.md#DowngradeLicense) | **Post** /api/licenses/downgrade | Remove the current license, causing the server to fall back to the default configuration.
[**GetLicense**](LicensesApi.md#GetLicense) | **Get** /api/licenses/{id} | Retrieve the license with the given id
[**ListLicenses**](LicensesApi.md#ListLicenses) | **Get** /api/licenses/ | Retrieve licenses
[**UpdateLicense**](LicensesApi.md#UpdateLicense) | **Put** /api/licenses/{id} | Update an existing license



## DowngradeLicense

> DowngradeLicense(ctx).Execute()

Remove the current license, causing the server to fall back to the default configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensesApi.DowngradeLicense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.DowngradeLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDowngradeLicenseRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> License GetLicense(ctx, id).Execute()

Retrieve the license with the given id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the license

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensesApi.GetLicense(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.GetLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the license | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**License**](License.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLicenses

> []License ListLicenses(ctx).Execute()

Retrieve licenses

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensesApi.ListLicenses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.ListLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLicenses`: []License
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.ListLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLicensesRequest struct via the builder pattern


### Return type

[**[]License**](License.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> License UpdateLicense(ctx, id).License(license).Execute()

Update an existing license

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the license
    license := *openapiclient.NewLicense("LicenseText_example") // License |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensesApi.UpdateLicense(context.Background(), id).License(license).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.UpdateLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.UpdateLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the license | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **license** | [**License**](License.md) |  | 

### Return type

[**License**](License.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

