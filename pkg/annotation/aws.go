package annotation

// Keep the YAML documentation format as it is used to render in the CRD public documentation. You can use Markdown in the documentation field.

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 14.0.0
//
// documentation:
//
//	This annotation allows configuration of the MINIMUM_IP_TARGET parameter for AWS CNI.
//	See [CNI Configuration Variables](https://github.com/aws/amazon-vpc-cni-k8s#cni-configuration-variables)
//	and [ENI and IP Target](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/docs/eni-and-ip-target.md)
const AWSCNIMinimumIPTarget = "alpha.cni.aws.giantswarm.io/minimum-ip-target"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.0
//
// documentation:
//
//	This annotation allows configuration of the ENABLE_PREFIX_DELEGATION parameter for AWS CNI.
//	See [Enable Prefix Delegation](https://github.com/aws/amazon-vpc-cni-k8s#enable_prefix_delegation-v190)
const AWSCNIPrefixDelegation = "alpha.cni.aws.giantswarm.io/prefix-delegation"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 14.0.0
//
// documentation:
//
//	This annotation allows configuration of the WARM_IP_TARGET parameter for AWS CNI.
//	See [CNI Configuration Variables](https://github.com/aws/amazon-vpc-cni-k8s#cni-configuration-variables)
//	and [ENI and IP Target](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/docs/eni-and-ip-target.md)
const AWSCNIWarmIPTarget = "alpha.cni.aws.giantswarm.io/warm-ip-target"

// support:
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 17.2.0
//
// documentation:
//
//	This annotation enables alpha feature EBS Volume Iops.
const AWSEBSVolumeIops = "alpha.aws.giantswarm.io/ebs-volume-iops"

// support:
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 17.2.0
//
// documentation:
//
//	This annotation enables alpha feature EBS Volume Throughput.
const AWSEBSVolumeThroughput = "alpha.aws.giantswarm.io/ebs-volume-throughput"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 17.1.0
//
// documentation:
//
//	This annotation enables alpha feature IAM Roles for Service Accounts.
//	See [IAM Roles for Service Accounts](https://docs.giantswarm.io/guides/iam-roles-for-service-accounts/)
const AWSIRSA = "alpha.aws.giantswarm.io/iam-roles-for-service-accounts"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//
// documentation:
//
//	This annotation is used for configuring maximum batch size for instances termination during ASG update.
//	The value can be either a whole number specifying the number of instances
//	or a percentage of total instances as decimal number ie `0.3` for 30%.
//	See [Fine Tuning Upgrades](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/)
//	and [AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize) for additional information.
const AWSUpdateMaxBatchSize = "alpha.aws.giantswarm.io/update-max-batch-size"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//
// documentation:
//
//	This annotation is used for configuring time pause between rolling a single batch during ASG update.
//	The value must be in ISO 8601 duration format, e. g. "PT5M" for five minutes or "PT10S" for 10 seconds.
//	See [Fine Tuning Upgrades](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/)
//	and [AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize) for additional information.
const AWSUpdatePauseTime = "alpha.aws.giantswarm.io/update-pause-time"

// Not documented as it's not usable.
// AWSMetadataV2 configures token usage for your AWS EC2 instance metadata requests.
// If the value is 'optional', you can choose to retrieve instance metadata with or without a signed token
// header on your request. If you retrieve the IAM role credentials without a token, the version 1.0 role
// credentials are returned. If you retrieve the IAM role credentials using a valid signed token, the version
// 2.0 role credentials are returned.
// If the state is 'required', you must send a signed token header with any instance metadata retrieval
// requests. In this state, retrieving the IAM role credentials always returns the version 2.0 credentials; the
// version 1.0 credentials are not available.
// Default value is 'optional'
//
//	[AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens)
const AWSMetadataV2 = "alpha.aws.giantswarm.io/metadata-v2"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 12.7.0
//
// documentation:
//
//	This annotation is used for configuring the subnet size of AWSCluster or AWSMachineDeployment.
//	The value is a number that will represent the subnet mask used when creating the subnet. This value must be smaller than 28 due to AWS restrictions.
const AWSSubnetSize = "alpha.aws.giantswarm.io/aws-subnet-size"

// support:
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 17.3.2
//
// documentation:
//
//	This annotation is used to support internal load balancers.
//	It will set a tag on the subnet of AWSMachineDeployment.
//	See [Subnet Discovery](https://github.com/kubernetes-sigs/aws-load-balancer-controller/blob/main/docs/deploy/subnet_discovery.md#private-subnets)
const AWSInternalELB = "alpha.aws.giantswarm.io/internal-elb"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//
// documentation:
//
//	This annotation is used to start a EC2 instance refresh in a Auto Scaling group.
//	See [aws-rolling-node-operator](https://github.com/giantswarm/aws-rolling-node-operator#operator-for-rolling-nodes)
const AWSInstanceRefresh = "alpha.aws.giantswarm.io/instance-refresh"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//
// documentation:
//
//	This annotation is used to cancel a EC2 instance refresh in a Auto Scaling group.
//	See [aws-rolling-node-operator](https://github.com/giantswarm/aws-rolling-node-operator#operator-for-rolling-nodes)
const AWSCancelInstanceRefresh = "alpha.aws.giantswarm.io/cancel-instance-refresh"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//
// documentation:
//
//	This annotation is used to set the amount of capacity which must remain healthy inside the Auto Scaling group.
//	See [aws-rolling-node-operator](https://github.com/giantswarm/aws-rolling-node-operator#operator-for-rolling-nodes)
const AWSInstanceRefreshMinHealthyPercentage = "alpha.aws.giantswarm.io/instance-refresh-min-healthy-percentage"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 16.1.1
//
// documentation:
//
//	This annotation is used to set the instance warmup which is the time period from when a new instance's state changes to InService to when it can receive traffic.
//	See [aws-rolling-node-operator](https://github.com/giantswarm/aws-rolling-node-operator#operator-for-rolling-nodes)
const AWSInstanceWarmupSeconds = "alpha.aws.giantswarm.io/instance-warmup-seconds"

// support:
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 19.0.0
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 19.0.0
//
// documentation:
//
//	This annotation is used to set the size of the logging volume in the nodes. The default is 100 (Gb).
const AWSLoggingVolumeSize = "alpha.aws.giantswarm.io/logging-volume-size"

// support:
//   - crd: awscontrolplanes.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 19.0.0
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 19.0.0
//
// documentation:
//
//	This annotation is used to set the size of the docker volume in the nodes. The default is value will be dockerVolumeSizeGB set in the CR.
const AWSContainerdVolumeSize = "alpha.aws.giantswarm.io/containerd-volume-size"
