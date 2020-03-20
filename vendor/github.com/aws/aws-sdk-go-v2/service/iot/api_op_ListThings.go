// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the ListThings operation.
type ListThingsInput struct {
	_ struct{} `type:"structure"`

	// The attribute name used to search for things.
	AttributeName *string `location:"querystring" locationName:"attributeName" type:"string"`

	// The attribute value used to search for things.
	AttributeValue *string `location:"querystring" locationName:"attributeValue" type:"string"`

	// The maximum number of results to return in this operation.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the thing type used to search for things.
	ThingTypeName *string `location:"querystring" locationName:"thingTypeName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListThingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AttributeName != nil {
		v := *s.AttributeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "attributeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AttributeValue != nil {
		v := *s.AttributeValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "attributeValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the ListThings operation.
type ListThingsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The things.
	Things []ThingAttribute `locationName:"things" type:"list"`
}

// String returns the string representation
func (s ListThingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Things != nil {
		v := s.Things

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "things", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListThings = "ListThings"

// ListThingsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists your things. Use the attributeName and attributeValue parameters to
// filter your things. For example, calling ListThings with attributeName=Color
// and attributeValue=Red retrieves all things in the registry that contain
// an attribute Color with the value Red.
//
//    // Example sending a request using ListThingsRequest.
//    req := client.ListThingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingsRequest(input *ListThingsInput) ListThingsRequest {
	op := &aws.Operation{
		Name:       opListThings,
		HTTPMethod: "GET",
		HTTPPath:   "/things",
	}

	if input == nil {
		input = &ListThingsInput{}
	}

	req := c.newRequest(op, input, &ListThingsOutput{})
	return ListThingsRequest{Request: req, Input: input, Copy: c.ListThingsRequest}
}

// ListThingsRequest is the request type for the
// ListThings API operation.
type ListThingsRequest struct {
	*aws.Request
	Input *ListThingsInput
	Copy  func(*ListThingsInput) ListThingsRequest
}

// Send marshals and sends the ListThings API request.
func (r ListThingsRequest) Send(ctx context.Context) (*ListThingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingsResponse{
		ListThingsOutput: r.Request.Data.(*ListThingsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingsResponse is the response type for the
// ListThings API operation.
type ListThingsResponse struct {
	*ListThingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThings request.
func (r *ListThingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
