// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartImageScanInput struct {
	_ struct{} `type:"structure"`

	// An object with identifying information for an Amazon ECR image.
	//
	// ImageId is a required field
	ImageId *ImageIdentifier `locationName:"imageId" type:"structure" required:"true"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to start an image scan request. If you do not specify a registry,
	// the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository that contains the images to scan.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s StartImageScanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartImageScanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartImageScanInput"}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageId != nil {
		if err := s.ImageId.Validate(); err != nil {
			invalidParams.AddNested("ImageId", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartImageScanOutput struct {
	_ struct{} `type:"structure"`

	// An object with identifying information for an Amazon ECR image.
	ImageId *ImageIdentifier `locationName:"imageId" type:"structure"`

	// The current state of the scan.
	ImageScanStatus *ImageScanStatus `locationName:"imageScanStatus" type:"structure"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s StartImageScanOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartImageScan = "StartImageScan"

// StartImageScanRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Starts an image vulnerability scan. An image scan can only be started once
// per day on an individual image. This limit includes if an image was scanned
// on initial push. For more information, see Image Scanning (https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html)
// in the Amazon Elastic Container Registry User Guide.
//
//    // Example sending a request using StartImageScanRequest.
//    req := client.StartImageScanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/StartImageScan
func (c *Client) StartImageScanRequest(input *StartImageScanInput) StartImageScanRequest {
	op := &aws.Operation{
		Name:       opStartImageScan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartImageScanInput{}
	}

	req := c.newRequest(op, input, &StartImageScanOutput{})
	return StartImageScanRequest{Request: req, Input: input, Copy: c.StartImageScanRequest}
}

// StartImageScanRequest is the request type for the
// StartImageScan API operation.
type StartImageScanRequest struct {
	*aws.Request
	Input *StartImageScanInput
	Copy  func(*StartImageScanInput) StartImageScanRequest
}

// Send marshals and sends the StartImageScan API request.
func (r StartImageScanRequest) Send(ctx context.Context) (*StartImageScanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartImageScanResponse{
		StartImageScanOutput: r.Request.Data.(*StartImageScanOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartImageScanResponse is the response type for the
// StartImageScan API operation.
type StartImageScanResponse struct {
	*StartImageScanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartImageScan request.
func (r *StartImageScanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
