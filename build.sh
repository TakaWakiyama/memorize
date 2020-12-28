build() {
  sam build
}
deploy () {
  sam deploy
}

startlocal() {
  cd container && docker-compose up -d
  cd ../
  sam local start-api --docker-network lambda-local
}

main() {
  build
  if [[ $1 = deploy ]];then
    deploy
  else
    startlocal
  fi
}

main $1