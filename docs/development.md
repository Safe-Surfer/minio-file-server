- [Preparations](#sec-1)
- [Install](#sec-2)
- [Accessing](#sec-3)
- [Uninstall](#sec-4)


# Preparations<a id="sec-1"></a>

Add the Helm repo:

```shell
helm repo add minio https://helm.min.io/
```

Create a development namespace:

```shell
kubectl create ns minio-file-server-dev
```

# Install<a id="sec-2"></a>

```shell
helm install minio \
  -n minio-file-server-dev \
  --set accessKey="minio-file-server"  \
  --set secretKey="minio-file-server" \
  --set defaultBucket.enabled=true \
  --set defaultBucket.name=minio \
  --set resources.requests.memory=1Gi \
  minio/minio
```

# Accessing<a id="sec-3"></a>

```shell
kubectl -n minio-file-server-dev port-forward svc/minio 9000
```

# Uninstall<a id="sec-4"></a>

```shell
helm uninstall minio -n minio-file-server-dev
```
