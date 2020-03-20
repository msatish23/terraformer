// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLoggingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the web ACL from which you want to delete
	// the LoggingConfiguration.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLoggingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoggingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLoggingConfigurationInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLoggingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLoggingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLoggingConfiguration = "DeleteLoggingConfiguration"

// DeleteLoggingConfigurationRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Permanently deletes the LoggingConfiguration from the specified web ACL.
//
//    // Example sending a request using DeleteLoggingConfigurationRequest.
//    req := client.DeleteLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteLoggingConfiguration
func (c *Client) DeleteLoggingConfigurationRequest(input *DeleteLoggingConfigurationInput) DeleteLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteLoggingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteLoggingConfigurationOutput{})
	return DeleteLoggingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteLoggingConfigurationRequest}
}

// DeleteLoggingConfigurationRequest is the request type for the
// DeleteLoggingConfiguration API operation.
type DeleteLoggingConfigurationRequest struct {
	*aws.Request
	Input *DeleteLoggingConfigurationInput
	Copy  func(*DeleteLoggingConfigurationInput) DeleteLoggingConfigurationRequest
}

// Send marshals and sends the DeleteLoggingConfiguration API request.
func (r DeleteLoggingConfigurationRequest) Send(ctx context.Context) (*DeleteLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoggingConfigurationResponse{
		DeleteLoggingConfigurationOutput: r.Request.Data.(*DeleteLoggingConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoggingConfigurationResponse is the response type for the
// DeleteLoggingConfiguration API operation.
type DeleteLoggingConfigurationResponse struct {
	*DeleteLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoggingConfiguration request.
func (r *DeleteLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
