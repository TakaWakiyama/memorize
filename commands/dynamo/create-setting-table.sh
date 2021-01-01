  aws dynamodb create-table \
  --table-name Settings \
  --attribute-definitions \
    AttributeName=user,AttributeType=S \
    AttributeName=category,AttributeType=S \
  --key-schema AttributeName=user,KeyType=HASH AttributeName=category,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=25,WriteCapacityUnits=1 \

