package ses

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// ConfigurationSetEventDestination_EventDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.EventDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html
type ConfigurationSetEventDestination_EventDestination struct {

	// CloudWatchDestination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-cloudwatchdestination
	CloudWatchDestination *ConfigurationSetEventDestination_CloudWatchDestination `json:"CloudWatchDestination,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-enabled
	Enabled *types.Value `json:"Enabled,omitempty"`

	// KinesisFirehoseDestination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-kinesisfirehosedestination
	KinesisFirehoseDestination *ConfigurationSetEventDestination_KinesisFirehoseDestination `json:"KinesisFirehoseDestination,omitempty"`

	// MatchingEventTypes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-matchingeventtypes
	MatchingEventTypes *types.Value `json:"MatchingEventTypes,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-name
	Name *types.Value `json:"Name,omitempty"`

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
func (r *ConfigurationSetEventDestination_EventDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.EventDestination"
}
