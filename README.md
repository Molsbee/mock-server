# mock-server
General purpose mock server that can be configured through rest api

Collection:
    Directory: CollectionName
        Group: File containing common routes

Example
Collection: BMC
    Groups:
        servers
        machines

Path: 
    /bmc/servers/{path}
    /bmc/machines/{path}


Route
```json
{
  "method": "GET",
  "path": "/",
  "headers": {
    "Authorization": "Bearer token",
    "Content-Type": "application/json"
  },
  "statusCode": 200,
  "body": "{\"some\": \"info\"}"
}
```
