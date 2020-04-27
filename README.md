# Link For You

## REST API Documentation

### Link management

- [Create new link](doc/link_post.md) : `POST /link/create`
- [Get link information](doc/link_get.md) : `GET /link/info/:domain/:hash`
- [Delete link](doc/link_delete.md) : `DELETE /link/delete/:domain/:hash`
- [Edit link parameters](doc/link_patch.md) : `PATCH /link/patch`
  
## Troubleshooting

- Count number of open connections on sqlite database

    ```bash
    lsof dbName.db
    ```
