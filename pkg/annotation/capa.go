package annotation

// Keep the YAML documentation format as it is used to render in the CRD public documentation. You can use Markdown in the documentation field.

// AWSVPCMode is set on AWSCluster. Possible values are AWSVPCModePublic and AWSVPCModePrivate
const AWSVPCMode = "aws.giantswarm.io/vpc-mode"

// AWSVPCModePublic is one possible AWSVPCMode annotation value. When it is set,
// CAPA controllers will create public-facing VPC for AWS cluster.
const AWSVPCModePublic = "public"

// AWSVPCModePrivate is one possible AWSVPCMode annotation value. When it is set,
// aws-vpc-operator will create private VPC for AWS cluster.
const AWSVPCModePrivate = "private"

// AWSDNSMode is set on AWSCluster.
const AWSDNSMode = "aws.giantswarm.io/dns-mode"

// AWSDNSAdditionalVPC is set on AWSCluster.
const AWSDNSAdditionalVPC = "aws.giantswarm.io/dns-assign-additional-vpc"

// DNSModePrivate is possible AWSDNSMode annotation value.
const DNSModePrivate = "private"
