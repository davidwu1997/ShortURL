# Short URL Service
## Local Run
``make deploy-local``

## Local Test

### Upload
![創建](./test/upload.png)
### Delete
![刪除](./test/delete.png)
### Redirect After Delete
![導向](./test/redirect.png)

## Tech Stack
* Gin
* wire
* Gorm
* go-redis

### RateLimit 每小時每個 ip 最多能請求 60 次