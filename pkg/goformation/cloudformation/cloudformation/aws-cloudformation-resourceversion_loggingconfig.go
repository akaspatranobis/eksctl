package cloudformation

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// ResourceVersion_LoggingConfig AWS CloudFormation Resource (AWS::CloudFormation::ResourceVersion.LoggingConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html
type ResourceVersion_LoggingConfig struct {

	// LogGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html#cfn-cloudformation-resourceversion-loggingconfig-loggroupname
	LogGroupName *types.Value `json:"LogGroupName,omitempty"`

	// LogRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html#cfn-cloudformation-resourceversion-loggingconfig-logrolearn
	LogRoleArn *types.Value `json:"LogRoleArn,omitempty"`

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
func (r *ResourceVersion_LoggingConfig) AWSCloudFormationType() string {
	return "AWS::CloudFormation::ResourceVersion.LoggingConfig"
}
