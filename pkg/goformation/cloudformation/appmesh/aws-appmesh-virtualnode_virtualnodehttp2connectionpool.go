package appmesh

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// VirtualNode_VirtualNodeHttp2ConnectionPool AWS CloudFormation Resource (AWS::AppMesh::VirtualNode.VirtualNodeHttp2ConnectionPool)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodehttp2connectionpool.html
type VirtualNode_VirtualNodeHttp2ConnectionPool struct {

	// MaxRequests AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodehttp2connectionpool.html#cfn-appmesh-virtualnode-virtualnodehttp2connectionpool-maxrequests
	MaxRequests *types.Value `json:"MaxRequests"`

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
func (r *VirtualNode_VirtualNodeHttp2ConnectionPool) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualNode.VirtualNodeHttp2ConnectionPool"
}
