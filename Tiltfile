#!/bin/python

yaml = helm(
  'deployments/minio-file-server',
  name='minio-file-server-dev',
  namespace='minio-file-server-dev',
  set=[
      "minio.host=minio:9000",
      "minio.bucket=bucket",
      "minio.SSL.enable=false",
      "minio.accessKey=minio-file-server",
      "minio.secretKey=minio-file-server",
      "service.type=NodePort"
  ]
  )
k8s_yaml(yaml)
load('ext://helm_remote', 'helm_remote')
helm_remote(
  'minio',
  repo_url="https://helm.min.io/",
  namespace='minio-file-server-dev',
  version="7.2.1",
  set=[
      "accessKey=minio-file-server",
      "secretKey=minio-file-server",
      "defaultBucket.enabled=true",
      "resources.requests.memory=1Gi"
  ]
  )

containerRepo='registry.gitlab.com/safesurfer/minio-file-server'
if os.getenv('KIND_EXPERIMENTAL_PROVIDER') == 'podman' and k8s_context() == 'kind-kind':
    custom_build(containerRepo, 'podman build -f build/Dockerfile -t $EXPECTED_REF . && podman save $EXPECTED_REF > /tmp/tilt-containerbuild.tar.gz && kind load image-archive /tmp/tilt-containerbuild.tar.gz', ['.'], disable_push=True, skips_local_docker=True)
else:
    docker_build(containerRepo, '.', dockerfile="build/Dockerfile")
allow_k8s_contexts('in-cluster')
