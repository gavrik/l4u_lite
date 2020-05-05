# Create new link

**URL** : `/link/create`  
**Method** : `POST`  
**Auth required** : YES  
**Data constraints**

- longLink - Full link to the resource. **Required**
- shortLink - Short link. If missed application generate it
- isEnabled - true | false
  
**Data example**

```json
{
    "longLink": "https://example.com",
    "shortLink": "execom",
    "isEnabled": true
}
```

## Success Response

___

**Condition** : If link created  
**Code** : `201 CREATED`  
**Content Example**

```json
{
    "longLink": "test-link_1",
    "shortLink": "CWI62a",
    "domain": "",
    "isEnabled": false,
    "creationOn": 1588259164
}
```

## Error Response

___

**Condition** : If not authorized
**Code** : 401 UNAUTHORIZED
**Content Example**

```json
{
    "error_number": 1,
    "error_message": "NotAuthorized"
}
```

**Condition** : If link not created
**Code** : 404 NOT FOUND
**Content Example**

```json
{
    "error_number": 5,
    "error_message": " [ Error description ] "
}
```

## Request example

___

```bash
curl -v --data '{"longLink": "test-link_1"}' --header "Authorization: TOKEN <admin token hash>" http://127.0.0.1:8081/link/create

*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8081 (#0)
> POST /link/create HTTP/1.1
> Host: 127.0.0.1:8081
> User-Agent: curl/7.64.1
> Accept: */*
> Authorization: TOKEN e319e2a5-95f5-48fd-bbc9-7315df21c382
> Content-Length: 27
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 27 out of 27 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json; charset=utf-8
< Date: Thu, 30 Apr 2020 15:15:24 GMT
< Content-Length: 101
<
* Connection #0 to host 127.0.0.1 left intact
{"longLink":"test-link_1","shortLink":"N7LYjh","domain":"","isEnabled":false,"creationOn":1588259724}* Closing connection 0

```
