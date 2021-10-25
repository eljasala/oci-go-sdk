// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListScriptsRequest wrapper for the ListScripts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListScripts.go.html to see an example of how to use ListScriptsRequest.
type ListScriptsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the content type given.
	ContentType *string `mandatory:"false" contributesTo:"query" name:"contentType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListScriptsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName and contentType is ascending.
	// Default order of timeCreated and timeUpdated is descending.
	// The displayName sort by is case insensitive.
	SortBy ListScriptsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScriptsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScriptsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScriptsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScriptsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListScriptsResponse wrapper for the ListScripts operation
type ListScriptsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScriptCollection instances
	ScriptCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScriptsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScriptsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScriptsSortOrderEnum Enum with underlying type: string
type ListScriptsSortOrderEnum string

// Set of constants representing the allowable values for ListScriptsSortOrderEnum
const (
	ListScriptsSortOrderAsc  ListScriptsSortOrderEnum = "ASC"
	ListScriptsSortOrderDesc ListScriptsSortOrderEnum = "DESC"
)

var mappingListScriptsSortOrder = map[string]ListScriptsSortOrderEnum{
	"ASC":  ListScriptsSortOrderAsc,
	"DESC": ListScriptsSortOrderDesc,
}

// GetListScriptsSortOrderEnumValues Enumerates the set of values for ListScriptsSortOrderEnum
func GetListScriptsSortOrderEnumValues() []ListScriptsSortOrderEnum {
	values := make([]ListScriptsSortOrderEnum, 0)
	for _, v := range mappingListScriptsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListScriptsSortByEnum Enum with underlying type: string
type ListScriptsSortByEnum string

// Set of constants representing the allowable values for ListScriptsSortByEnum
const (
	ListScriptsSortByDisplayname ListScriptsSortByEnum = "displayName"
	ListScriptsSortByTimecreated ListScriptsSortByEnum = "timeCreated"
	ListScriptsSortByTimeupdated ListScriptsSortByEnum = "timeUpdated"
	ListScriptsSortByContenttype ListScriptsSortByEnum = "contentType"
)

var mappingListScriptsSortBy = map[string]ListScriptsSortByEnum{
	"displayName": ListScriptsSortByDisplayname,
	"timeCreated": ListScriptsSortByTimecreated,
	"timeUpdated": ListScriptsSortByTimeupdated,
	"contentType": ListScriptsSortByContenttype,
}

// GetListScriptsSortByEnumValues Enumerates the set of values for ListScriptsSortByEnum
func GetListScriptsSortByEnumValues() []ListScriptsSortByEnum {
	values := make([]ListScriptsSortByEnum, 0)
	for _, v := range mappingListScriptsSortBy {
		values = append(values, v)
	}
	return values
}
