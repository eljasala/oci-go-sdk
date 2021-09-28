// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v48/common"
	"net/http"
)

// DeletePrivateEndpointRequest wrapper for the DeletePrivateEndpoint operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataflow/DeletePrivateEndpoint.go.html to see an example of how to use DeletePrivateEndpointRequest.
type DeletePrivateEndpointRequest struct {

	// The unique ID for a private endpoint.
	PrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"privateEndpointId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource,
	// set the `if-match` parameter to the value of the etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeletePrivateEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeletePrivateEndpointRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeletePrivateEndpointRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeletePrivateEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeletePrivateEndpointResponse wrapper for the DeletePrivateEndpoint operation
type DeletePrivateEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle assigned identifier for a work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response DeletePrivateEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeletePrivateEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
