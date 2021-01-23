# bin/bash
if [[ $1 == loc ]];then
  opt="--endpoint-url http://localhost:8000"
fi

aws dynamodb put-item --table-name Memos \
--item '{ "User": { "S": "Twaki" }, "date_created": { "S": "2021-01-02" }, "MemoId": { "S": "0004" }, "url": { "S": "https://www.google.com/search?q=kk&oq" },"MemoType": { "S": "url" }}' \
$opt
