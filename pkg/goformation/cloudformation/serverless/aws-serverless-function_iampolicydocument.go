package serverless

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Function_IAMPolicyDocument AWS CloudFormation Resource (AWS::Serverless::Function.IAMPolicyDocument)
// See: http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
type Function_IAMPolicyDocument struct {

	// Statement AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
	Statement interface{} `json:"Statement,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_IAMPolicyDocument) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.IAMPolicyDocument"
}
