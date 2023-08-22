# Reciept: a take home challenge for https://github.com/fetch-rewards/receipt-processor-challenge
------------

to run locally: 'go run main.go'

cURL for proccessReciept endpoint

curl --location 'http://localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
  ],
  "total": "35.35"
}'




cURL for getPoints endpoint

curl --location 'http://localhost:8080/receipts/7c020a77-f9b5-4458-aa0f-21cbbb1bc7ad/points' \
