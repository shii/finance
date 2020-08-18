

# Run API on Postman

### Tranfer ###
```
curl -X POST \
  http://localhost:8080/v1/transfer \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 83bebf7a-1fe3-c032-7a89-cb53fcc4cbcc' \
  -d '{
"account_from": 1234,
"account_to": 7890,
"amount" : 200
}'
```
### Balance ###
```
curl -X GET \
  'http://localhost:8080/v1/balance?id=1234' \
  -H 'cache-control: no-cache' \
  -H 'postman-token: c1d3d329-cd55-afd8-6f2f-4a8a7f9cd3c3'
  ````
