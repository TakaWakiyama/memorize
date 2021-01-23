# bin/bash

aws events put-rule \
  --name test_cli \
  --schedule-expression "cron(0 20 * * ? *)" \
  --state DISABLED \
  --cli-input-json "{}"

aws events put-targets --rule test_cli \
  --targets "Id"="testpyt","Arn"="arn:aws:lambda:ap-northeast-1:255282616662:function:testpyt","Input"='"{}"'