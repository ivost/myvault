#!/bin/bash
#set -euo pipefail

#echo === health
#url=localhost:8080/myvault
# http $url/health
url=localhost:8080
http $url/health

# kubectl get mappings -A
# kubectl describe mapping myvault-map1
# kubectl describe mapping myvault-map2

# this doesn't work unless local
#grpc=myvault.VaultService
#grpcurl -plaintext 192.168.99.100:30506  ${grpc}/Health

# but this is OK
../build/client  -config client-config.yaml

# elvis is OK too (something with reflection?)

#echo === REST endpoint on 8080/myvault
#curl -X POST  http://localhost:8080/hello -d '{ "myvaulting": { "first_name": "John" } }'
#
#echo === grpcurl
#
#grpcurl -plaintext localhost:52052 describe myredis.KVService
#
#grpcurl -d '{ "myvaulting": { "first_name": "Ivo" } }' \
#     -plaintext localhost:52052  myredis.KVService/myvault
#
#echo === grpc go client
#[[ -f build/client ]] && build/client
#
#
#echo " "
#
#echo === evans CLI
#
#echo '{ "myvaulting": { "first_name": "Ivo" } }' | evans --port 52052 --package myvault --service KVService --call myvault myvault/myvault.proto
#
#echo === evans REPL
#
#echo evans myvault/myvault.proto --repl --host localhost --port 52052 -r
#
#echo evans myvault/myvault.proto --repl --host 192.168.99.100 --port 30506 -r
#
#echo package myvault
#echo service KVService
#echo call myvault

#  pod=$(kubectl get pod -l app=hello -o  jsonpath='{.items[0].metadata.name}')
#  kubectl port-forward $pod 52052:52052 8080:8080

# amb.
# http://192.168.99.100:30506/myvault

# url=http://192.168.99.100:30506/myvault
# http $url/health
# curl -X POST $url/hello -d '{ "myvaulting": { "first_name": "John", "last_name": "Doe" } }'

# grpcurl -plaintext 192.168.99.100:30506 myredis.KVService/Health
