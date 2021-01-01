# bin/bash

aws dynamodb get-item --table-name Items  --key '{ "user": {"S": "Twaki" },"item_id": {"S": "0001" }  }'