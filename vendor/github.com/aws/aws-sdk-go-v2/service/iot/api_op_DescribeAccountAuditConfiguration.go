// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeAccountAuditConfigurationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountAuditConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAccountAuditConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type DescribeAccountAuditConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Which audit checks are enabled and disabled for this account.
	AuditCheckConfigurations map[string]AuditCheckConfiguration `locationName:"auditCheckConfigurations" type:"map"`

	// Information about the targets to which audit notifications are sent for this
	// account.
	AuditNotificationTargetConfigurations map[string]AuditNotificationTarget `locationName:"auditNotificationTargetConfigurations" type:"map"`

	// The ARN of the role that grants permission to AWS IoT to access information
	// about your devices, policies, certificates, and other items as required when
	// performing an audit.
	//
	// On the first call to UpdateAccountAuditConfiguration, this parameter is required.
	RoleArn *string `locationName:"roleArn" min:"20" type:"string"`
}

// String returns the string representation
func (s DescribeAccountAuditConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAccountAuditConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuditCheckConfigurations != nil {
		v := s.AuditCheckConfigurations

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "auditCheckConfigurations", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.AuditNotificationTargetConfigurations != nil {
		v := s.AuditNotificationTargetConfigurations

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "auditNotificationTargetConfigurations", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeAccountAuditConfiguration = "DescribeAccountAuditConfiguration"

// DescribeAccountAuditConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks
// are enabled or disabled.
//
//    // Example sending a request using DescribeAccountAuditConfigurationRequest.
//    req := client.DescribeAccountAuditConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeAccountAuditConfigurationRequest(input *DescribeAccountAuditConfigurationInput) DescribeAccountAuditConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAuditConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/audit/configuration",
	}

	if input == nil {
		input = &DescribeAccountAuditConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountAuditConfigurationOutput{})
	return DescribeAccountAuditConfigurationRequest{Request: req, Input: input, Copy: c.DescribeAccountAuditConfigurationRequest}
}

// DescribeAccountAuditConfigurationRequest is the request type for the
// DescribeAccountAuditConfiguration API operation.
type DescribeAccountAuditConfigurationRequest struct {
	*aws.Request
	Input *DescribeAccountAuditConfigurationInput
	Copy  func(*DescribeAccountAuditConfigurationInput) DescribeAccountAuditConfigurationRequest
}

// Send marshals and sends the DescribeAccountAuditConfiguration API request.
func (r DescribeAccountAuditConfigurationRequest) Send(ctx context.Context) (*DescribeAccountAuditConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAuditConfigurationResponse{
		DescribeAccountAuditConfigurationOutput: r.Request.Data.(*DescribeAccountAuditConfigurationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAuditConfigurationResponse is the response type for the
// DescribeAccountAuditConfiguration API operation.
type DescribeAccountAuditConfigurationResponse struct {
	*DescribeAccountAuditConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAuditConfiguration request.
func (r *DescribeAccountAuditConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
