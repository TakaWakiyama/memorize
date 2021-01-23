if [[ $1 == loc ]];then
  opt="--endpoint-url http://localhost:8000"
fi

aws dynamodb delete-item --table-name Memos \
  --key '{"User":{"S":"Twaki"},"MemoId":{"S":"0004"}}' \
  $opt
