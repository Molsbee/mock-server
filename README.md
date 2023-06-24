# mock-server
General purpose mock server that can be configured through rest api


Collection
    Name: servers

Api Endpoint Prefixed by Collection
/servers

servers
List of Routes related to the collection
{
    Method: "GET",
    StatusCode: 200,
    Body: "{\"data\": \"content\"}",
    ContentType: "application/json",
    ResponseHeaders: [
        "Authorization: Bearer Token",
        "Content-Type: application/json"
    ]
}