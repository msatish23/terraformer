// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ValidateSecurityProfileBehaviorsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the behaviors that, when violated by a device (thing), cause an
	// alert.
	//
	// Behaviors is a required field
	Behaviors []Behavior `locationName:"behaviors" type:"list" required:"true"`
}

// String returns the string representation
func (s ValidateSecurityProfileBehaviorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValidateSecurityProfileBehaviorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValidateSecurityProfileBehaviorsInput"}

	if s.Behaviors == nil {
		invalidParams.Add(aws.NewErrParamRequired("Behaviors"))
	}
	if s.Behaviors != nil {
		for i, v := range s.Behaviors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Behaviors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ValidateSecurityProfileBehaviorsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Behaviors != nil {
		v := s.Behaviors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "behaviors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type ValidateSecurityProfileBehaviorsOutput struct {
	_ struct{} `type:"structure"`

	// True if the behaviors were valid.
	Valid *bool `locationName:"valid" type:"boolean"`

	// The list of any errors found in the behaviors.
	ValidationErrors []ValidationError `locationName:"validationErrors" type:"list"`
}

// String returns the string representation
func (s ValidateSecurityProfileBehaviorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ValidateSecurityProfileBehaviorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Valid != nil {
		v := *s.Valid

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "valid", protocol.BoolValue(v), metadata)
	}
	if s.ValidationErrors != nil {
		v := s.ValidationErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "validationErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opValidateSecurityProfileBehaviors = "ValidateSecurityProfileBehaviors"

// ValidateSecurityProfileBehaviorsRequest returns a request value for making API operation for
// AWS IoT.
//
// Validates a Device Defender security profile behaviors specification.
//
//    // Example sending a request using ValidateSecurityProfileBehaviorsRequest.
//    req := client.ValidateSecurityProfileBehaviorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ValidateSecurityProfileBehaviorsRequest(input *ValidateSecurityProfileBehaviorsInput) ValidateSecurityProfileBehaviorsRequest {
	op := &aws.Operation{
		Name:       opValidateSecurityProfileBehaviors,
		HTTPMethod: "POST",
		HTTPPath:   "/security-profile-behaviors/validate",
	}

	if input == nil {
		input = &ValidateSecurityProfileBehaviorsInput{}
	}

	req := c.newRequest(op, input, &ValidateSecurityProfileBehaviorsOutput{})
	return ValidateSecurityProfileBehaviorsRequest{Request: req, Input: input, Copy: c.ValidateSecurityProfileBehaviorsRequest}
}

// ValidateSecurityProfileBehaviorsRequest is the request type for the
// ValidateSecurityProfileBehaviors API operation.
type ValidateSecurityProfileBehaviorsRequest struct {
	*aws.Request
	Input *ValidateSecurityProfileBehaviorsInput
	Copy  func(*ValidateSecurityProfileBehaviorsInput) ValidateSecurityProfileBehaviorsRequest
}

// Send marshals and sends the ValidateSecurityProfileBehaviors API request.
func (r ValidateSecurityProfileBehaviorsRequest) Send(ctx context.Context) (*ValidateSecurityProfileBehaviorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidateSecurityProfileBehaviorsResponse{
		ValidateSecurityProfileBehaviorsOutput: r.Request.Data.(*ValidateSecurityProfileBehaviorsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidateSecurityProfileBehaviorsResponse is the response type for the
// ValidateSecurityProfileBehaviors API operation.
type ValidateSecurityProfileBehaviorsResponse struct {
	*ValidateSecurityProfileBehaviorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidateSecurityProfileBehaviors request.
func (r *ValidateSecurityProfileBehaviorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
