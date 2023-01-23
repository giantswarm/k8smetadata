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

// NetworkTopologyModeAnnotation is the annotation indicating the network topology mode a cluster uses
// Valid values are "GiantSwarmManaged", "CustomerManaged" and "None"
const NetworkTopologyModeAnnotation = "network-topology.giantswarm.io/mode"

const NetworkTopologyModeGiantSwarmManaged = "GiantSwarmManaged"
const NetworkTopologyModeUserManaged = "UserManaged"
const NetworkTopologyModeNone = "None"

// NetworkTopologyTransitGatewayIDAnnotation contains the ID of the Transit Gateway used by the cluster.
// This is either the user-provided TGW or the one created by this operator.
const NetworkTopologyTransitGatewayIDAnnotation = "network-topology.giantswarm.io/transit-gateway"

// NetworkTopologyPrefixListIDAnnotation contains the ID of the Prefix List containing the CIDRs of all clusters.
// This is either the user-provided PL ID or the one created by this operator.
const NetworkTopologyPrefixListIDAnnotation = "network-topology.giantswarm.io/prefix-list"

// ResolverRulesOwnerAWSAccountId contains the AWS Account ID that owns the resolver rules that need to be associated
// with the workload cluster VPC.
const ResolverRulesOwnerAWSAccountId = "aws.giantswarm.io/resolver-rules-owner-account"
