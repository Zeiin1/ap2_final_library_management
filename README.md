# Order Microservice
Create `dev.env` file in `pkg/config/envs` __(create directory envs)__, then paste the code:

```env
PORT=:8081
DB_URL=postgres://{username}:{password}@{db_host}:5432/services
BOOK_SVC_URL={svc_host}:8082
// svc_host = localhost -> on local machine (pc)
// svc_host = book-svc -> on Docker
// db_host = localhost -> on local machine (pc)
// db_host = postgresql -> on Docker
```

Exchange the values in _{} brackets_ with needed.

---

It will not create a database `services` by itself. I ask you to create it by own!

---

To run the Order Microservice, write it in terminal:
```
make server
```

---
