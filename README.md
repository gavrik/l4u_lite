# Link For You

Count number of open connection in an sql database

```bash
lsof dbName.db
```

## API

POST: /link/create  
{"longLink":"long_link_url","shortLink":"short_link_string","domain":"","isEnabled":true|false}

GET: /link/info/:domain/:hash  

PATCH: /link/patch  
{"longLink":"long_link_url","shortLink":"short_link_string","domain":"","isEnabled":true|false}

DELETE: /link/delete/:domain/:hash  
