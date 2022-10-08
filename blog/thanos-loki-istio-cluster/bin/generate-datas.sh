#!/bin/sh

CONTEXT_NAME=$1
CLUSTER_NAME=$2

SAVE_PATH=/tmp/${CLUSTER_NAME}/
TAR_PATH=/tmp/prom-${CLUSTER_NAME}.tar.gz

# Get S3 secret
kubectl --context $CONTEXT_NAME get secret cluster-objstore-secret -n monitoring -o jsonpath='{.data.objstore\.yml}' | base64 -d > /tmp/a

export S3_ENDPOINT_URL="http://`cat /tmp/a | dasel -r yaml ".config.endpoint"`"
export S3_BUCKET=`cat /tmp/a | dasel -r yaml '.config.bucket'`
export S3_BUCKET=`cat /tmp/a | dasel -r yaml '.config.bucket'`
export AWS_ACCESS_KEY_ID=`cat /tmp/a | dasel -r yaml '.config.access_key'`
export AWS_SECRET_ACCESS_KEY=`cat /tmp/a | dasel -r yaml '.config.secret_key'`

# generate blocks
thanosbench block plan -p continuous-365d-tiny --labels 'cluster="${CLUSTER_NAME}"' --labels 'namespace="test"' --labels 'app="fsdfds"' | thanosbench block gen --output.dir=$SAVE_PATH

# tar
CURRENT=`pwd`
cd $SAVE_PATH && tar -cvzf $TAR_PATH .
cd $CURRENT

rm -rf chunks_head

echo "Start upload to S3"
# Upload to s3
s5cmd cp $TAR_PATH s3://${S3_BUCKET}/

