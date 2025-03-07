# \AccelerationApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccelerationShow**](AccelerationApi.md#AccelerationShow) | **Get** /domains/{domain}/acceleration | Get the content of acceleration settings
[**AccelerationUpdate**](AccelerationApi.md#AccelerationUpdate) | **Patch** /domains/{domain}/acceleration | Update the content of acceleration settings
[**ImageResizeShow**](AccelerationApi.md#ImageResizeShow) | **Get** /domains/{domain}/image-resize | Get the content of image resize settings
[**ImageResizeUpdate**](AccelerationApi.md#ImageResizeUpdate) | **Patch** /domains/{domain}/image-resize | Update the content of image resize settings



## AccelerationShow

> AccelerationResponse AccelerationShow(ctx, domain).Execute()

Get the content of acceleration settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arash-r1c/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationApi.AccelerationShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationApi.AccelerationShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccelerationShow`: AccelerationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationApi.AccelerationShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccelerationShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccelerationResponse**](AccelerationResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## AccelerationUpdate

> AccelerationResponse AccelerationUpdate(ctx, domain).AccelerationUpdate(accelerationUpdate).Execute()

Update the content of acceleration settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arash-r1c/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    accelerationUpdate := *openapiclient.NewAccelerationUpdate() // AccelerationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationApi.AccelerationUpdate(context.Background(), domain).AccelerationUpdate(accelerationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationApi.AccelerationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccelerationUpdate`: AccelerationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationApi.AccelerationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccelerationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accelerationUpdate** | [**AccelerationUpdate**](AccelerationUpdate.md) |  | 

### Return type

[**AccelerationResponse**](AccelerationResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ImageResizeShow

> ImageResizeResponse ImageResizeShow(ctx, domain).Execute()

Get the content of image resize settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arash-r1c/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationApi.ImageResizeShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationApi.ImageResizeShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageResizeShow`: ImageResizeResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationApi.ImageResizeShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageResizeShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageResizeResponse**](ImageResizeResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ImageResizeUpdate

> ImageResizeResponse ImageResizeUpdate(ctx, domain).ImageResize(imageResize).Execute()

Update the content of image resize settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arash-r1c/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    imageResize := *openapiclient.NewImageResize() // ImageResize |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationApi.ImageResizeUpdate(context.Background(), domain).ImageResize(imageResize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationApi.ImageResizeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageResizeUpdate`: ImageResizeResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationApi.ImageResizeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageResizeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageResize** | [**ImageResize**](ImageResize.md) |  | 

### Return type

[**ImageResizeResponse**](ImageResizeResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

