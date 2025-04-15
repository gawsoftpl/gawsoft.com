TOKEN=`curl -X 'POST' \
  'http://localhost:3000/token' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "token_name": "test",
  "data": {
    "user": "test",
    "aud": "test",
    "sub": "cluster"
  }
}' | jq -r '.token'`
CONTROL_PLANE=`docker port kind-control-plane 6443`

curl -k https://$CONTROL_PLANE/api \
  -H "Authorization: Bearer $TOKEN"