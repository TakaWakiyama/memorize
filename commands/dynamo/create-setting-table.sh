  aws dynamodb create-table \
  --table-name Settings \
  --attribute-definitions \
    AttributeName=user,AttributeType=S \
    AttributeName=setting_id,AttributeType=S \
  --key-schema AttributeName=user,KeyType=HASH AttributeName=setting_id,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=25,WriteCapacityUnits=1 \

