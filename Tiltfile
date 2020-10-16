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
docker_build('registry.gitlab.com/safesurfer/minio-file-server', '.', dockerfile="build/Dockerfile")
allow_k8s_contexts('in-cluster')
