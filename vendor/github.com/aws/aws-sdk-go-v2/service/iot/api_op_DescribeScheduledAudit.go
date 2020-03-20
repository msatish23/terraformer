// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The name of the scheduled audit whose information you want to get.
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScheduledAuditInput"}

	if s.ScheduledAuditName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledAuditName"))
	}
	if s.ScheduledAuditName != nil && len(*s.ScheduledAuditName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledAuditName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeScheduledAuditInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ScheduledAuditName != nil {
		v := *s.ScheduledAuditName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "scheduledAuditName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeScheduledAuditOutput struct {
	_ struct{} `type:"structure"`

	// The day of the month on which the scheduled audit takes place. Will be "1"
	// through "31" or "LAST". If days 29-31 are specified, and the month does not
	// have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `locationName:"dayOfMonth" type:"string"`

	// The day of the week on which the scheduled audit takes place. One of "SUN",
	// "MON", "TUE", "WED", "THU", "FRI", or "SAT".
	DayOfWeek DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// How often the scheduled audit takes place. One of "DAILY", "WEEKLY", "BIWEEKLY",
	// or "MONTHLY". The start time of each audit is determined by the system.
	Frequency AuditFrequency `locationName:"frequency" type:"string" enum:"true"`

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string `locationName:"scheduledAuditArn" type:"string"`

	// The name of the scheduled audit.
	ScheduledAuditName *string `locationName:"scheduledAuditName" min:"1" type:"string"`

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list
	// of all checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list"`
}

// String returns the string representation
func (s DescribeScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeScheduledAuditOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DayOfMonth != nil {
		v := *s.DayOfMonth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfMonth", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DayOfWeek) > 0 {
		v := s.DayOfWeek

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfWeek", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Frequency) > 0 {
		v := s.Frequency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "frequency", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ScheduledAuditArn != nil {
		v := *s.ScheduledAuditArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "scheduledAuditArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ScheduledAuditName != nil {
		v := *s.ScheduledAuditName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "scheduledAuditName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetCheckNames != nil {
		v := s.TargetCheckNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetCheckNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opDescribeScheduledAudit = "DescribeScheduledAudit"

// DescribeScheduledAuditRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about a scheduled audit.
//
//    // Example sending a request using DescribeScheduledAuditRequest.
//    req := client.DescribeScheduledAuditRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeScheduledAuditRequest(input *DescribeScheduledAuditInput) DescribeScheduledAuditRequest {
	op := &aws.Operation{
		Name:       opDescribeScheduledAudit,
		HTTPMethod: "GET",
		HTTPPath:   "/audit/scheduledaudits/{scheduledAuditName}",
	}

	if input == nil {
		input = &DescribeScheduledAuditInput{}
	}

	req := c.newRequest(op, input, &DescribeScheduledAuditOutput{})
	return DescribeScheduledAuditRequest{Request: req, Input: input, Copy: c.DescribeScheduledAuditRequest}
}

// DescribeScheduledAuditRequest is the request type for the
// DescribeScheduledAudit API operation.
type DescribeScheduledAuditRequest struct {
	*aws.Request
	Input *DescribeScheduledAuditInput
	Copy  func(*DescribeScheduledAuditInput) DescribeScheduledAuditRequest
}

// Send marshals and sends the DescribeScheduledAudit API request.
func (r DescribeScheduledAuditRequest) Send(ctx context.Context) (*DescribeScheduledAuditResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScheduledAuditResponse{
		DescribeScheduledAuditOutput: r.Request.Data.(*DescribeScheduledAuditOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScheduledAuditResponse is the response type for the
// DescribeScheduledAudit API operation.
type DescribeScheduledAuditResponse struct {
	*DescribeScheduledAuditOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScheduledAudit request.
func (r *DescribeScheduledAuditResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
