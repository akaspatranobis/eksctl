package wafv2

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// WebACL_CustomRequestHandling AWS CloudFormation Resource (AWS::WAFv2::WebACL.CustomRequestHandling)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html
type WebACL_CustomRequestHandling struct {

	// InsertHeaders AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html#cfn-wafv2-webacl-customrequesthandling-insertheaders
	InsertHeaders []WebACL_CustomHTTPHeader `json:"InsertHeaders,omitempty"`

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
func (r *WebACL_CustomRequestHandling) AWSCloudFormationType() string {
	return "AWS::WAFv2::WebACL.CustomRequestHandling"
}
