# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.19.0] - 2023-01-23

### Added

- Add `aws.giantswarm.io/resolver-rules-owner-account` annotation.

## [0.18.0] - 2023-01-13

### Added

- Add `alpha.aws.giantswarm.io/logging-volume-size` annotation.
- Add `alpha.aws.giantswarm.io/containerd-volume-size` annotation.

## [0.17.0] - 2023-01-09

### Added

- Add `aws.giantswarm.io/dns-mode` annotation.
- Add `aws.giantswarm.io/dns-assign-additional-vpc` annotation.
- Add network topology annotations.

## [0.16.1] - 2022-12-02

### Fixed

- Annotation for `aws-rolling-node-operator` needs start with `alpha.aws`.

## [0.16.0] - 2022-11-30

### Added

- Add instance warmup seconds annotation.

## [0.15.0] - 2022-10-10

### Added

- Add `aws.giantswarm.io/vpc-mode` annotation.

### Changed

- Update go to v1.19.

## [0.14.0] - 2022-10-06

### Added

- Add `cilium.giantswarm.io/force-disable-cilium-kube-proxy` annotation.

## [0.13.0] - 2022-09-20

## Added

- Add AWS instance refresh annotations.

## [0.12.0] - 2022-08-02

### Added

- Add `cilium.giantswarm.io/pod-cidr` annotation.

## [0.11.1] - 2022-06-03

### Added

- Add `giantswarm.io/service-priority` label.

## [0.11.0] - 2022-05-05

### Added

- Add `alpha.giantswarm.io/internal-elb` annotation.

## [0.10.1] - 2022-03-21

### Fixed

- Support `alpha.aws.giantswarm.io/ebs-volume-throughput` in AWS v17.2.0.
- Support `alpha.aws.giantswarm.io/ebs-volume-iops` in AWS v17.2.0.

## [0.10.0] - 2022-03-16

### Added

- Add `alpha.aws.giantswarm.io/ebs-volume-throughput` annotation.
- Add `alpha.aws.giantswarm.io/ebs-volume-iops` annotation.

## [0.9.4] - 2022-03-04

### Added

- Add `alpha.aws.giantswarm.io/iam-roles-for-service-accounts` annotation.

## [0.9.3] - 2022-02-16

### Added

- Add annotation `app-operator.giantswarm.io/trigger-reconciliation`.

## [0.9.2] - 2022-02-01

### Added

- Add annotation `node.giantswarm.io/cgroupv1`.

## [0.9.1] - 2022-01-21

## [0.9.0] - 2022-01-21

### Added

- Add annotations for the encryption-provider-operator.

## [0.8.0] - 2022-01-18

### Added

- Add `AppName` annotation for use by chart CR watcher in app-operator.

## [0.7.1] - 2021-12-08

### Fixed

- Move `ClusterAppsOperatorWatching` to `label` package instead of `annotation` as intended.

## [0.7.0] - 2021-12-08

### Added

- Add `ClusterAppsOperatorWatching` label as a replacement for `ClusterAppsOperatorVersion` when
  only one version of the operator is running to indicate which clusters should be watched (in
  cases where `cluster-operator` is running in the same cluster).

## [0.6.0] - 2021-11-16

### Added

- Add AWSCNIPrefixDelegation annotation.

## [0.5.0] - 2021-11-11

### Added

- Add annotation to indicate a resource shouldn't be removed (`giantswarm.io/keep`)

## [0.4.0] - 2021-10-04

### Added

- Add annotations for `upgrade-schedule-operator`

## [0.3.0] - 2021-05-12

- Add ChartOperatorForceHelmUpggrade annotation.
- Add ClusterAppsOperatorVersion label.

## [0.2.0] - 2021-05-11

### Added

- Add annotations from https://github.com/giantswarm/app/tree/v4.12.0/pkg/annotation
- Add annotations from https://github.com/giantswarm/app-operator/tree/v4.4.0/pkg/annotation
- Add labels from https://github.com/giantswarm/app/tree/v4.12.0/pkg/label

## [0.1.0] - 2021-04-23

### Added

- Add annotations from https://github.com/giantswarm/apiextensions/tree/v3.22.0/pkg/annotation
- Add labels from https://github.com/giantswarm/apiextensions/tree/v3.22.0/pkg/label

[Unreleased]: https://github.com/giantswarm/k8smetadata/compare/v0.19.0...HEAD
[0.19.0]: https://github.com/giantswarm/k8smetadata/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/giantswarm/k8smetadata/compare/v0.17.0...v0.18.0
[0.17.0]: https://github.com/giantswarm/k8smetadata/compare/v0.16.1...v0.17.0
[0.16.1]: https://github.com/giantswarm/k8smetadata/compare/v0.16.0...v0.16.1
[0.16.0]: https://github.com/giantswarm/k8smetadata/compare/v0.15.0...v0.16.0
[0.15.0]: https://github.com/giantswarm/k8smetadata/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/giantswarm/k8smetadata/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/giantswarm/k8smetadata/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/giantswarm/k8smetadata/compare/v0.11.1...v0.12.0
[0.11.1]: https://github.com/giantswarm/k8smetadata/compare/v0.11.0...v0.11.1
[0.11.0]: https://github.com/giantswarm/k8smetadata/compare/v0.10.1...v0.11.0
[0.10.1]: https://github.com/giantswarm/k8smetadata/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/giantswarm/k8smetadata/compare/v0.9.4...v0.10.0
[0.9.4]: https://github.com/giantswarm/k8smetadata/compare/v0.9.3...v0.9.4
[0.9.3]: https://github.com/giantswarm/k8smetadata/compare/v0.9.2...v0.9.3
[0.9.2]: https://github.com/giantswarm/k8smetadata/compare/v0.9.1...v0.9.2
[0.9.1]: https://github.com/giantswarm/k8smetadata/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/giantswarm/k8smetadata/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/k8smetadata/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/k8smetadata/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/k8smetadata/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/k8smetadata/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/k8smetadata/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/k8smetadata/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/k8smetadata/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/k8smetadata/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/k8smetadata/releases/tag/v0.1.0
