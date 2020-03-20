// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DescribeCACertificate operation.
type DescribeCACertificateInput struct {
	_ struct{} `type:"structure"`

	// The CA certificate identifier.
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"caCertificateId" min:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCACertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCACertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCACertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCACertificateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "caCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the DescribeCACertificate operation.
type DescribeCACertificateOutput struct {
	_ struct{} `type:"structure"`

	// The CA certificate description.
	CertificateDescription *CACertificateDescription `locationName:"certificateDescription" type:"structure"`

	// Information about the registration configuration.
	RegistrationConfig *RegistrationConfig `locationName:"registrationConfig" type:"structure"`
}

// String returns the string representation
func (s DescribeCACertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCACertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateDescription != nil {
		v := s.CertificateDescription

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "certificateDescription", v, metadata)
	}
	if s.RegistrationConfig != nil {
		v := s.RegistrationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "registrationConfig", v, metadata)
	}
	return nil
}

const opDescribeCACertificate = "DescribeCACertificate"

// DescribeCACertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes a registered CA certificate.
//
//    // Example sending a request using DescribeCACertificateRequest.
//    req := client.DescribeCACertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeCACertificateRequest(input *DescribeCACertificateInput) DescribeCACertificateRequest {
	op := &aws.Operation{
		Name:       opDescribeCACertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/cacertificate/{caCertificateId}",
	}

	if input == nil {
		input = &DescribeCACertificateInput{}
	}

	req := c.newRequest(op, input, &DescribeCACertificateOutput{})
	return DescribeCACertificateRequest{Request: req, Input: input, Copy: c.DescribeCACertificateRequest}
}

// DescribeCACertificateRequest is the request type for the
// DescribeCACertificate API operation.
type DescribeCACertificateRequest struct {
	*aws.Request
	Input *DescribeCACertificateInput
	Copy  func(*DescribeCACertificateInput) DescribeCACertificateRequest
}

// Send marshals and sends the DescribeCACertificate API request.
func (r DescribeCACertificateRequest) Send(ctx context.Context) (*DescribeCACertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCACertificateResponse{
		DescribeCACertificateOutput: r.Request.Data.(*DescribeCACertificateOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCACertificateResponse is the response type for the
// DescribeCACertificate API operation.
type DescribeCACertificateResponse struct {
	*DescribeCACertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCACertificate request.
func (r *DescribeCACertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
