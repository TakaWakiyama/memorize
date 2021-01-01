# bin/bash
aws dynamodb create-table \
  --table-name Items \
  --attribute-definitions \
    AttributeName=user,AttributeType=S \
    AttributeName=item_id,AttributeType=S \
    AttributeName=date_created,AttributeType=S \
    AttributeName=item_type,AttributeType=S \
  --key-schema AttributeName=user,KeyType=HASH AttributeName=item_id,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --local-secondary-indexes \
    IndexName=date_created,KeySchema=[{'AttributeName=user,KeyType=HASH'},{'AttributeName=date_created,KeyType=RANGE'}],Projection={ProjectionType=ALL} \
    IndexName=item_type,KeySchema=[{'AttributeName=user,KeyType=HASH'},{'AttributeName=item_type,KeyType=RANGE'}],Projection={ProjectionType=ALL}
