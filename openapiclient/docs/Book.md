# Book

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Edition** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 

## Methods

### NewBook

`func NewBook() *Book`

NewBook instantiates a new Book object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookWithDefaults

`func NewBookWithDefaults() *Book`

NewBookWithDefaults instantiates a new Book object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Book) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Book) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Book) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Book) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Book) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Book) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Book) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Book) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAuthor

`func (o *Book) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Book) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Book) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Book) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetIsbn

`func (o *Book) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *Book) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *Book) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *Book) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetYear

`func (o *Book) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Book) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Book) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Book) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetEdition

`func (o *Book) GetEdition() int32`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *Book) GetEditionOk() (*int32, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *Book) SetEdition(v int32)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *Book) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetRating

`func (o *Book) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Book) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Book) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Book) HasRating() bool`

HasRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


