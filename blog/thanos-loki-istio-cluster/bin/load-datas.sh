#!/bin/sh

CONTEXT_NAME=$1
CLUSTER_NAME=$2

# Get S3 secret
kubectl --context $CONTEXT_NAME get secret cluster-objstore-secret -n monitoring -o jsonpath='{.data.objstore\.yml}' | base64 -d > /tmp/a

S3_ENDPOINT_URL="http://`cat /tmp/a | dasel -r yaml ".config.endpoint"`"
S3_BUCKET=`cat /tmp/a | dasel -r yaml '.config.bucket'`
AWS_ACCESS_KEY_ID=`cat /tmp/a | dasel -r yaml '.config.access_key'`
AWS_SECRET_ACCESS_KEY=`cat /tmp/a | dasel -r yaml '.config.secret_key'`

cat <<EOF > /tmp/script
wget -O s5cmd.tar.gz https://github.com/peak/s5cmd/releases/download/v2.0.0/s5cmd_2.0.0_Linux-64bit.tar.gz
tar -xvf s5cmd.tar.gz
chmod +x s5cmd
export S3_ENDPOINT_URL=$S3_ENDPOINT_URL
export S3_BUCKET=$S3_BUCKET
export AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
./s5cmd cp s3://${S3_BUCKET}/prom-${CLUSTER_NAME}.tar.gz .
tar -xvf prom-${CLUSTER_NAME}.tar.gz
EOF

SCRIPT=`cat /tmp/script`
echo $SCRIPT

kubectl --context $CONTEXT_NAME  exec -it prometheus-prometheus-kube-prometheus-prometheus-0 -n monitoring -- /bin/sh -c "$SCRIPT"

