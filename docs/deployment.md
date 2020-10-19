- [Helm](#sec-1)
  - [Preliminary steps](#sec-1-1)
  - [Installation](#sec-1-2)
  - [Upgrading versions](#sec-1-3)
  - [Uninstalling](#sec-1-4)


# Helm<a id="sec-1"></a>

## Preliminary steps<a id="sec-1-1"></a>

Create a namespace:

```shell
kubectl create ns minio-file-server
```

## Installation<a id="sec-1-2"></a>

Install with Helm:

```shell
helm install minio-file-server-dev \
  -n minio-file-server-dev \
  --set minio.accessKey="minio-file-server" \
  --set minio.secretKey="minio-file-server" \
deployments/minio-file-server
```

Note: to configure, please check out [the configuration docs](./configuration.md)

## Upgrading versions<a id="sec-1-3"></a>

Upgrade a release with Helm:

```shell
helm upgrade minio-file-server-dev \
  -n minio-file-server-dev \
  --set minio.accessKey="minio-file-server" \
  --set minio.secretKey="minio-file-server" \
deployments/minio-file-server
```

## Uninstalling<a id="sec-1-4"></a>

Uninstall with Helm:

```shell
helm uninstall minio-file-server-dev \
  -n minio-file-server-dev
```
