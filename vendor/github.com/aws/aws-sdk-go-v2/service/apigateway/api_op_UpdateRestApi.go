// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to update an existing RestApi resource in your collection.
type UpdateRestApiInput struct {
	_ struct{} `type:"structure"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRestApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRestApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRestApiInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRestApiInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a REST API.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type UpdateRestApiOutput struct {
	_ struct{} `type:"structure"`

	// The source of the API key for metering requests according to a usage plan.
	// Valid values are:
	//    * HEADER to read the API key from the X-API-Key header of a request.
	//
	//    * AUTHORIZER to read the API key from the UsageIdentifierKey from a custom
	//    authorizer.
	ApiKeySource ApiKeySourceType `locationName:"apiKeySource" type:"string" enum:"true"`

	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string `locationName:"binaryMediaTypes" type:"list"`

	// The timestamp when the API was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// The API's description.
	Description *string `locationName:"description" type:"string"`

	// The endpoint configuration of this RestApi showing the endpoint types of
	// the API.
	EndpointConfiguration *EndpointConfiguration `locationName:"endpointConfiguration" type:"structure"`

	// The API's identifier. This identifier is unique across all of your APIs in
	// API Gateway.
	Id *string `locationName:"id" type:"string"`

	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with
	// a null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int64 `locationName:"minimumCompressionSize" type:"integer"`

	// The API's name.
	Name *string `locationName:"name" type:"string"`

	// A stringified JSON policy document that applies to this RestApi regardless
	// of the caller and Method configuration.
	Policy *string `locationName:"policy" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// A version identifier for the API.
	Version *string `locationName:"version" type:"string"`

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s UpdateRestApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRestApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ApiKeySource) > 0 {
		v := s.ApiKeySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySource", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.BinaryMediaTypes != nil {
		v := s.BinaryMediaTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "binaryMediaTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointConfiguration != nil {
		v := s.EndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "endpointConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MinimumCompressionSize != nil {
		v := *s.MinimumCompressionSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumCompressionSize", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opUpdateRestApi = "UpdateRestApi"

// UpdateRestApiRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about the specified API.
//
//    // Example sending a request using UpdateRestApiRequest.
//    req := client.UpdateRestApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateRestApiRequest(input *UpdateRestApiInput) UpdateRestApiRequest {
	op := &aws.Operation{
		Name:       opUpdateRestApi,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}",
	}

	if input == nil {
		input = &UpdateRestApiInput{}
	}

	req := c.newRequest(op, input, &UpdateRestApiOutput{})
	return UpdateRestApiRequest{Request: req, Input: input, Copy: c.UpdateRestApiRequest}
}

// UpdateRestApiRequest is the request type for the
// UpdateRestApi API operation.
type UpdateRestApiRequest struct {
	*aws.Request
	Input *UpdateRestApiInput
	Copy  func(*UpdateRestApiInput) UpdateRestApiRequest
}

// Send marshals and sends the UpdateRestApi API request.
func (r UpdateRestApiRequest) Send(ctx context.Context) (*UpdateRestApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRestApiResponse{
		UpdateRestApiOutput: r.Request.Data.(*UpdateRestApiOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRestApiResponse is the response type for the
// UpdateRestApi API operation.
type UpdateRestApiResponse struct {
	*UpdateRestApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRestApi request.
func (r *UpdateRestApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
