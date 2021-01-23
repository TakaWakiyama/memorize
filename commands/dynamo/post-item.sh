# bin/bash
if [[ $1 == loc ]];then
  opt="--endpoint-url http://localhost:8000"
fi

aws dynamodb put-item --table-name Memos \
  --item file://fixtures/items.json \
  $opt
