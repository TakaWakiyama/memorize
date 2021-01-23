# bin/bash
if [[ $1 == loc ]];then
  opt="--endpoint-url http://localhost:8000"
fi

aws dynamodb query --table-name Memos \
--index-name User-MemoType-index \
--key-condition-expression "#User = :v_user and MemoType = :v_memo_type" \
--expression-attribute-values  '{":v_user":{"S": "Twaki"}, ":v_memo_type":{"S":"url"} }' \
--expression-attribute-names '{"#User":"User"}' \
$opt
