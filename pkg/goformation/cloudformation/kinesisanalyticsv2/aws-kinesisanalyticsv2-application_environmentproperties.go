package kinesisanalyticsv2

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Application_EnvironmentProperties AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.EnvironmentProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html
type Application_EnvironmentProperties struct {

	// PropertyGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html#cfn-kinesisanalyticsv2-application-environmentproperties-propertygroups
	PropertyGroups []Application_PropertyGroup `json:"PropertyGroups,omitempty"`

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
func (r *Application_EnvironmentProperties) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.EnvironmentProperties"
}
