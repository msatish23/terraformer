// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input to the UpdateCACertificate operation.
type UpdateCACertificateInput struct {
	_ struct{} `type:"structure"`

	// The CA certificate identifier.
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"caCertificateId" min:"64" type:"string" required:"true"`

	// The new value for the auto registration status. Valid values are: "ENABLE"
	// or "DISABLE".
	NewAutoRegistrationStatus AutoRegistrationStatus `location:"querystring" locationName:"newAutoRegistrationStatus" type:"string" enum:"true"`

	// The updated status of the CA certificate.
	//
	// Note: The status value REGISTER_INACTIVE is deprecated and should not be
	// used.
	NewStatus CACertificateStatus `location:"querystring" locationName:"newStatus" type:"string" enum:"true"`

	// Information about the registration configuration.
	RegistrationConfig *RegistrationConfig `locationName:"registrationConfig" type:"structure"`

	// If true, removes auto registration.
	RemoveAutoRegistration *bool `locationName:"removeAutoRegistration" type:"boolean"`
}

// String returns the string representation
func (s UpdateCACertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCACertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCACertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}
	if s.RegistrationConfig != nil {
		if err := s.RegistrationConfig.Validate(); err != nil {
			invalidParams.AddNested("RegistrationConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCACertificateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistrationConfig != nil {
		v := s.RegistrationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "registrationConfig", v, metadata)
	}
	if s.RemoveAutoRegistration != nil {
		v := *s.RemoveAutoRegistration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "removeAutoRegistration", protocol.BoolValue(v), metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "caCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.NewAutoRegistrationStatus) > 0 {
		v := s.NewAutoRegistrationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "newAutoRegistrationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.NewStatus) > 0 {
		v := s.NewStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "newStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type UpdateCACertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCACertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCACertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateCACertificate = "UpdateCACertificate"

// UpdateCACertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates a registered CA certificate.
//
//    // Example sending a request using UpdateCACertificateRequest.
//    req := client.UpdateCACertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateCACertificateRequest(input *UpdateCACertificateInput) UpdateCACertificateRequest {
	op := &aws.Operation{
		Name:       opUpdateCACertificate,
		HTTPMethod: "PUT",
		HTTPPath:   "/cacertificate/{caCertificateId}",
	}

	if input == nil {
		input = &UpdateCACertificateInput{}
	}

	req := c.newRequest(op, input, &UpdateCACertificateOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateCACertificateRequest{Request: req, Input: input, Copy: c.UpdateCACertificateRequest}
}

// UpdateCACertificateRequest is the request type for the
// UpdateCACertificate API operation.
type UpdateCACertificateRequest struct {
	*aws.Request
	Input *UpdateCACertificateInput
	Copy  func(*UpdateCACertificateInput) UpdateCACertificateRequest
}

// Send marshals and sends the UpdateCACertificate API request.
func (r UpdateCACertificateRequest) Send(ctx context.Context) (*UpdateCACertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCACertificateResponse{
		UpdateCACertificateOutput: r.Request.Data.(*UpdateCACertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCACertificateResponse is the response type for the
// UpdateCACertificate API operation.
type UpdateCACertificateResponse struct {
	*UpdateCACertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCACertificate request.
func (r *UpdateCACertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
