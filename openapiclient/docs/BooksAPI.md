# \BooksAPI

All URIs are relative to *https://api.bookstore.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBook**](BooksAPI.md#AddBook) | **Post** /books | Add a new book
[**DeleteBook**](BooksAPI.md#DeleteBook) | **Delete** /books/{bookId} | Delete a book by ID
[**GetBookById**](BooksAPI.md#GetBookById) | **Get** /books/{bookId} | Get a book by ID
[**ListBooks**](BooksAPI.md#ListBooks) | **Get** /books | List all books
[**UpdateBook**](BooksAPI.md#UpdateBook) | **Put** /books/{bookId} | Update a book by ID



## AddBook

> Book AddBook(ctx).CreateBook(createBook).Execute()

Add a new book

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createBook := *openapiclient.NewCreateBook() // CreateBook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.AddBook(context.Background()).CreateBook(createBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.AddBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBook`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.AddBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBook** | [**CreateBook**](CreateBook.md) |  | 

### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBook

> DeleteBook(ctx, bookId).Execute()

Delete a book by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "bookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BooksAPI.DeleteBook(context.Background(), bookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.DeleteBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookById

> Book GetBookById(ctx, bookId).Execute()

Get a book by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "bookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookById(context.Background(), bookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookById`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBooks

> []Book ListBooks(ctx).Execute()

List all books

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.ListBooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.ListBooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBooks`: []Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.ListBooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBooksRequest struct via the builder pattern


### Return type

[**[]Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBook

> Book UpdateBook(ctx, bookId).UpdateBook(updateBook).Execute()

Update a book by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "bookId_example" // string | 
	updateBook := *openapiclient.NewUpdateBook() // UpdateBook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.UpdateBook(context.Background(), bookId).UpdateBook(updateBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.UpdateBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBook`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.UpdateBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBook** | [**UpdateBook**](UpdateBook.md) |  | 

### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

