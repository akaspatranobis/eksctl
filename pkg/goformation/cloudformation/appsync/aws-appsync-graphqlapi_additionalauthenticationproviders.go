package appsync

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// GraphQLApi_AdditionalAuthenticationProviders AWS CloudFormation Resource (AWS::AppSync::GraphQLApi.AdditionalAuthenticationProviders)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-additionalauthenticationproviders.html
type GraphQLApi_AdditionalAuthenticationProviders struct {

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
func (r *GraphQLApi_AdditionalAuthenticationProviders) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLApi.AdditionalAuthenticationProviders"
}
