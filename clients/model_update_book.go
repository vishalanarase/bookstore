/*
Bookstore API

API for managing a bookstore.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateBook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBook{}

// UpdateBook struct for UpdateBook
type UpdateBook struct {
	Title *string `json:"title,omitempty"`
	Author *string `json:"author,omitempty"`
	Isbn *string `json:"isbn,omitempty"`
	Price *float32 `json:"price,omitempty"`
	PublishedDate *string `json:"published_date,omitempty"`
}

// NewUpdateBook instantiates a new UpdateBook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBook() *UpdateBook {
	this := UpdateBook{}
	return &this
}

// NewUpdateBookWithDefaults instantiates a new UpdateBook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBookWithDefaults() *UpdateBook {
	this := UpdateBook{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateBook) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBook) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateBook) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateBook) SetTitle(v string) {
	o.Title = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *UpdateBook) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBook) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *UpdateBook) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *UpdateBook) SetAuthor(v string) {
	o.Author = &v
}

// GetIsbn returns the Isbn field value if set, zero value otherwise.
func (o *UpdateBook) GetIsbn() string {
	if o == nil || IsNil(o.Isbn) {
		var ret string
		return ret
	}
	return *o.Isbn
}

// GetIsbnOk returns a tuple with the Isbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBook) GetIsbnOk() (*string, bool) {
	if o == nil || IsNil(o.Isbn) {
		return nil, false
	}
	return o.Isbn, true
}

// HasIsbn returns a boolean if a field has been set.
func (o *UpdateBook) HasIsbn() bool {
	if o != nil && !IsNil(o.Isbn) {
		return true
	}

	return false
}

// SetIsbn gets a reference to the given string and assigns it to the Isbn field.
func (o *UpdateBook) SetIsbn(v string) {
	o.Isbn = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UpdateBook) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBook) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UpdateBook) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *UpdateBook) SetPrice(v float32) {
	o.Price = &v
}

// GetPublishedDate returns the PublishedDate field value if set, zero value otherwise.
func (o *UpdateBook) GetPublishedDate() string {
	if o == nil || IsNil(o.PublishedDate) {
		var ret string
		return ret
	}
	return *o.PublishedDate
}

// GetPublishedDateOk returns a tuple with the PublishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBook) GetPublishedDateOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedDate) {
		return nil, false
	}
	return o.PublishedDate, true
}

// HasPublishedDate returns a boolean if a field has been set.
func (o *UpdateBook) HasPublishedDate() bool {
	if o != nil && !IsNil(o.PublishedDate) {
		return true
	}

	return false
}

// SetPublishedDate gets a reference to the given string and assigns it to the PublishedDate field.
func (o *UpdateBook) SetPublishedDate(v string) {
	o.PublishedDate = &v
}

func (o UpdateBook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Isbn) {
		toSerialize["isbn"] = o.Isbn
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PublishedDate) {
		toSerialize["published_date"] = o.PublishedDate
	}
	return toSerialize, nil
}

type NullableUpdateBook struct {
	value *UpdateBook
	isSet bool
}

func (v NullableUpdateBook) Get() *UpdateBook {
	return v.value
}

func (v *NullableUpdateBook) Set(val *UpdateBook) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBook) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBook(val *UpdateBook) *NullableUpdateBook {
	return &NullableUpdateBook{value: val, isSet: true}
}

func (v NullableUpdateBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


