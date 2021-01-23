cd ../.. && sam local invoke \
  -e commands/functions/events/notification.json \
  --docker-network lambda-local \
  --profile local \
  SendNotification