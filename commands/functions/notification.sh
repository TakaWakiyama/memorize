cd ../.. && sam local invoke \
  SendNotification \
  -e commands/functions/events/notification.json \
  --docker-network lambda-local \
  --profile local \
  $1
