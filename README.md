# Gwynbliedd Helm Charts Repository

## Overview

This repository contains Helm charts for various applications and services. These charts simplify the deployment of applications on Kubernetes clusters.

## Usage

### Adding the Repository

To add this Helm repository to your local Helm configuration:

`bash
helm repo add gwynbliedd https://gutsbutcher.github.io/helm-charts/
helm repo update
`

### Available Charts

| Chart Name | Version | App Version | Description | Installation Command |
|-----------|---------|-------------|-------------|----------------------|
| nginx | 0.1.0 | 1.24.0 | A helm chart for nginx deployment | `helm install nginx gwynbliedd/nginx` |
| test | 0.1.0 | 1.16.0 | A Helm chart for Kubernetes | `helm install test gwynbliedd/test` |
| test | 0.1.0 | 1.16.0 | A Helm chart for Kubernetes | `helm install test gwynbliedd/test` |
| webapp | 0.1.0 | 1.16.0 | A Helm chart for Kubernetes | `helm install webapp gwynbliedd/webapp` |

## Chart Development

### Prerequisites

- Kubernetes 1.20+
- Helm 3.x

### Installation

To install a specific chart:

`bash
helm install <release-name> gwynbliedd/<chart-name>
`

## Last Updated

*Automatically generated on 2025-03-25 14:23:47*

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create a new branch
3. Make your changes
4. Submit a pull request

## License

(Specify your license, e.g., MIT, Apache 2.0)

## Support

For issues or questions, please [open an issue](https://github.com/gutsbutcher/helm-charts/issues) in the GitHub repository.
