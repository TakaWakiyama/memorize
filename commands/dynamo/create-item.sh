# bin/bash

aws dynamodb put-item --table-name Items \
--item '{ "user": { "S": "Twaki" }, "date_created": { "S": "2021-01-02" }, "item_id": { "S": "0005" }, "url": { "S": "https://www.google.com/search?q=kk&oq" },"item_type": { "S": "url" }}'