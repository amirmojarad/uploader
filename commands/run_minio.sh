docker run \
  -p 9000:9000 \
  -p 9001:9001 \
  --name minio1 \
  -v ~/minio/data:/data \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  quay.io/minio/minio \
  server /data --console-address ":9001"

#docker run
#  -p 9000:9000
#  -p 9001:9001
#  -e "MINIO_ROOT_USER=..."
#  -e "MINIO_ROOT_PASSWORD=..."
#  quay.io/minio/minio
#  server /data --console-address ":9001"