// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateDeploymentGroup operation.
type CreateDeploymentGroupInput struct {
	_ struct{} `type:"structure"`

	// Information to add about Amazon CloudWatch alarms when the deployment group
	// is created.
	AlarmConfiguration *AlarmConfiguration `locationName:"alarmConfiguration" type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// Configuration information for an automatic rollback that is added when a
	// deployment group is created.
	AutoRollbackConfiguration *AutoRollbackConfiguration `locationName:"autoRollbackConfiguration" type:"structure"`

	// A list of associated Amazon EC2 Auto Scaling groups.
	AutoScalingGroups []string `locationName:"autoScalingGroups" type:"list"`

	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration *BlueGreenDeploymentConfiguration `locationName:"blueGreenDeploymentConfiguration" type:"structure"`

	// If specified, the deployment configuration name can be either one of the
	// predefined configurations provided with AWS CodeDeploy or a custom deployment
	// configuration that you create by calling the create deployment configuration
	// operation.
	//
	// CodeDeployDefault.OneAtATime is the default deployment configuration. It
	// is used if a configuration isn't specified for the deployment or deployment
	// group.
	//
	// For more information about the predefined deployment configurations in AWS
	// CodeDeploy, see Working with Deployment Groups in AWS CodeDeploy (https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html)
	// in the AWS CodeDeploy User Guide.
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string"`

	// The name of a new deployment group for the specified application.
	//
	// DeploymentGroupName is a required field
	DeploymentGroupName *string `locationName:"deploymentGroupName" min:"1" type:"string" required:"true"`

	// Information about the type of deployment, in-place or blue/green, that you
	// want to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle *DeploymentStyle `locationName:"deploymentStyle" type:"structure"`

	// The Amazon EC2 tags on which to filter. The deployment group includes EC2
	// instances with any of the specified tags. Cannot be used in the same call
	// as ec2TagSet.
	Ec2TagFilters []EC2TagFilter `locationName:"ec2TagFilters" type:"list"`

	// Information about groups of tags applied to EC2 instances. The deployment
	// group includes only EC2 instances identified by all the tag groups. Cannot
	// be used in the same call as ec2TagFilters.
	Ec2TagSet *EC2TagSet `locationName:"ec2TagSet" type:"structure"`

	// The target Amazon ECS services in the deployment group. This applies only
	// to deployment groups that use the Amazon ECS compute platform. A target Amazon
	// ECS service is specified as an Amazon ECS cluster and service name pair using
	// the format <clustername>:<servicename>.
	EcsServices []ECSService `locationName:"ecsServices" type:"list"`

	// Information about the load balancer used in a deployment.
	LoadBalancerInfo *LoadBalancerInfo `locationName:"loadBalancerInfo" type:"structure"`

	// The on-premises instance tags on which to filter. The deployment group includes
	// on-premises instances with any of the specified tags. Cannot be used in the
	// same call as OnPremisesTagSet.
	OnPremisesInstanceTagFilters []TagFilter `locationName:"onPremisesInstanceTagFilters" type:"list"`

	// Information about groups of tags applied to on-premises instances. The deployment
	// group includes only on-premises instances identified by all of the tag groups.
	// Cannot be used in the same call as onPremisesInstanceTagFilters.
	OnPremisesTagSet *OnPremisesTagSet `locationName:"onPremisesTagSet" type:"structure"`

	// A service role ARN that allows AWS CodeDeploy to act on the user's behalf
	// when interacting with AWS services.
	//
	// ServiceRoleArn is a required field
	ServiceRoleArn *string `locationName:"serviceRoleArn" type:"string" required:"true"`

	// The metadata that you apply to CodeDeploy deployment groups to help you organize
	// and categorize them. Each tag consists of a key and an optional value, both
	// of which you define.
	Tags []Tag `locationName:"tags" type:"list"`

	// Information about triggers to create when the deployment group is created.
	// For examples, see Create a Trigger for an AWS CodeDeploy Event (https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-sns.html)
	// in the AWS CodeDeploy User Guide.
	TriggerConfigurations []TriggerConfig `locationName:"triggerConfigurations" type:"list"`
}

// String returns the string representation
func (s CreateDeploymentGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentGroupInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}

	if s.DeploymentGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentGroupName"))
	}
	if s.DeploymentGroupName != nil && len(*s.DeploymentGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentGroupName", 1))
	}

	if s.ServiceRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDeploymentGroup operation.
type CreateDeploymentGroupOutput struct {
	_ struct{} `type:"structure"`

	// A unique deployment group ID.
	DeploymentGroupId *string `locationName:"deploymentGroupId" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDeploymentGroup = "CreateDeploymentGroup"

// CreateDeploymentGroupRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Creates a deployment group to which application revisions are deployed.
//
//    // Example sending a request using CreateDeploymentGroupRequest.
//    req := client.CreateDeploymentGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateDeploymentGroup
func (c *Client) CreateDeploymentGroupRequest(input *CreateDeploymentGroupInput) CreateDeploymentGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDeploymentGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDeploymentGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentGroupOutput{})
	return CreateDeploymentGroupRequest{Request: req, Input: input, Copy: c.CreateDeploymentGroupRequest}
}

// CreateDeploymentGroupRequest is the request type for the
// CreateDeploymentGroup API operation.
type CreateDeploymentGroupRequest struct {
	*aws.Request
	Input *CreateDeploymentGroupInput
	Copy  func(*CreateDeploymentGroupInput) CreateDeploymentGroupRequest
}

// Send marshals and sends the CreateDeploymentGroup API request.
func (r CreateDeploymentGroupRequest) Send(ctx context.Context) (*CreateDeploymentGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentGroupResponse{
		CreateDeploymentGroupOutput: r.Request.Data.(*CreateDeploymentGroupOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentGroupResponse is the response type for the
// CreateDeploymentGroup API operation.
type CreateDeploymentGroupResponse struct {
	*CreateDeploymentGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeploymentGroup request.
func (r *CreateDeploymentGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
