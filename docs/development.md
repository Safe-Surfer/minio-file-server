- [Preparations](#sec-1)
  - [Install and bring up Minikube or Kind](#sec-1-1)
    - [Kind](#sec-1-1-1)
    - [Minikube](#sec-1-1-2)
  - [Add the Helm repo](#sec-1-2)
  - [Create a development namespace](#sec-1-3)
- [Install](#sec-2)
- [Accessing](#sec-3)
- [Configuring a `.env` file](#sec-4)
- [Launching](#sec-5)
- [Uninstall](#sec-6)
- [Docs](#sec-7)


# Preparations<a id="sec-1"></a>

## Install and bring up Minikube or Kind<a id="sec-1-1"></a>

### Kind<a id="sec-1-1-1"></a>

Start the cluster:

```shell
kind create cluster
```

### Minikube<a id="sec-1-1-2"></a>

Start the cluster:

```shell
minikube start
```

## Add the Helm repo<a id="sec-1-2"></a>

```shell
helm repo add minio https://helm.min.io/
```

## Create a development namespace<a id="sec-1-3"></a>

```shell
kubectl create ns minio-file-server-dev
```

# Install<a id="sec-2"></a>

Install Minio into:

```shell
helm install minio \
  -n minio-file-server-dev \
  --set accessKey="minio-file-server"  \
  --set secretKey="minio-file-server" \
  --set defaultBucket.enabled=true \
  --set resources.requests.memory=1Gi \
  minio/minio
```

# Accessing<a id="sec-3"></a>

Bind `:9000` on your host to minio in the development cluster

```shell
kubectl -n minio-file-server-dev port-forward svc/minio 9000
```

# Configuring a `.env` file<a id="sec-4"></a>

Write a configuration file, something similar to this in `.env` at the root of the repo:

```shell
APP_MINIO_HOST=localhost:9000
APP_MINIO_ACCESS_KEY=minio-file-server
APP_MINIO_SECRET_KEY=minio-file-server
APP_MINIO_BUCKET=bucket
APP_MINIO_USE_SSL=false
```

# Launching<a id="sec-5"></a>

Compile and run:

```shell
go run main.go
```

# Uninstall<a id="sec-6"></a>

Remove minio from the development cluster

```shell
helm uninstall minio -n minio-file-server-dev
```

# Docs<a id="sec-7"></a>

To run the docs in development, use:

```sh
docker run --rm -it -p 8000:8000 -v ${PWD}:/docs:ro,Z squidfunk/mkdocs-material
```
