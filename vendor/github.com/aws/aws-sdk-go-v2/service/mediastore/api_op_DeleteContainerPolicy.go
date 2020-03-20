// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteContainerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that holds the policy.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteContainerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteContainerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteContainerPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteContainerPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteContainerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteContainerPolicy = "DeleteContainerPolicy"

// DeleteContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the access policy that is associated with the specified container.
//
//    // Example sending a request using DeleteContainerPolicyRequest.
//    req := client.DeleteContainerPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerPolicy
func (c *Client) DeleteContainerPolicyRequest(input *DeleteContainerPolicyInput) DeleteContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteContainerPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteContainerPolicyOutput{})
	return DeleteContainerPolicyRequest{Request: req, Input: input, Copy: c.DeleteContainerPolicyRequest}
}

// DeleteContainerPolicyRequest is the request type for the
// DeleteContainerPolicy API operation.
type DeleteContainerPolicyRequest struct {
	*aws.Request
	Input *DeleteContainerPolicyInput
	Copy  func(*DeleteContainerPolicyInput) DeleteContainerPolicyRequest
}

// Send marshals and sends the DeleteContainerPolicy API request.
func (r DeleteContainerPolicyRequest) Send(ctx context.Context) (*DeleteContainerPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteContainerPolicyResponse{
		DeleteContainerPolicyOutput: r.Request.Data.(*DeleteContainerPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteContainerPolicyResponse is the response type for the
// DeleteContainerPolicy API operation.
type DeleteContainerPolicyResponse struct {
	*DeleteContainerPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteContainerPolicy request.
func (r *DeleteContainerPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
