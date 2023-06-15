# Order Microservice
Create `dev.env` file in `pkg/config/envs` _(create directory envs)_, then paste the code:

```env
PORT=:8081
DB_URL=postgres://{username}:{password}@{host}:5432/services
BOOK_SVC_URL={svc_host}:8082
// svc_host = localhost -> on local machine (pc)
// svc_host = book-svc -> on Docker 
```

Exchange the values in _{} brackets_ with needed.

---

To run the Order Microservice, write it in terminal:
```
make server
```

---
