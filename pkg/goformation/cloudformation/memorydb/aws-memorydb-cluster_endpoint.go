package memorydb

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// Cluster_Endpoint AWS CloudFormation Resource (AWS::MemoryDB::Cluster.Endpoint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html
type Cluster_Endpoint struct {

	// Address AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html#cfn-memorydb-cluster-endpoint-address
	Address *types.Value `json:"Address,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-cluster-endpoint.html#cfn-memorydb-cluster-endpoint-port
	Port *types.Value `json:"Port,omitempty"`

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
func (r *Cluster_Endpoint) AWSCloudFormationType() string {
	return "AWS::MemoryDB::Cluster.Endpoint"
}
