yaml = helm(
  'deployments/minio-file-server',
  name='minio-file-server-dev',
  namespace='minio-file-server-dev',
  set=[
      "accessKey=minio-file-server",
      "secretKey=minio-file-server"
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
