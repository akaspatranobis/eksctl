package iot

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Authorizer_TokenSigningPublicKeys AWS CloudFormation Resource (AWS::IoT::Authorizer.TokenSigningPublicKeys)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-authorizer-tokensigningpublickeys.html
type Authorizer_TokenSigningPublicKeys struct {

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
func (r *Authorizer_TokenSigningPublicKeys) AWSCloudFormationType() string {
	return "AWS::IoT::Authorizer.TokenSigningPublicKeys"
}
