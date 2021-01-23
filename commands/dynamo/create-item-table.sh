# bin/bash
if [[ $1 != prod ]];then
  opt="--endpoint-url http://localhost:8000"
fi

aws dynamodb create-table \
  --cli-input-json file://fixtures/item-table-definition.json \
  $opt

