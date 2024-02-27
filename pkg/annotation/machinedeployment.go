package annotation

// MachineDeploymentName is the node pool annotation where human-friendly node pool
// name set by the customer is stored.
const MachineDeploymentName = "machine-deployment.giantswarm.io/name"

// MachineDeploymentSubnet is used to specify the IPv4 CIDR used for a
// node pool defined by a MachineDeployment or AWSMachineDeployment CR.
const MachineDeploymentSubnet = "machine-deployment.giantswarm.io/subnet"
