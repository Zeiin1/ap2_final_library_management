# Book Microservice
Create `dev.env` file in `pkg/config/envs` _(create directory envs)_, then paste the code:

```env
PORT=:8082
DB_URL=postgres://{username}:{password}@{host}:5432/services
```

Exchange the values in _{} brackets_ with needed.

---

To run the Book Microservice, write it in terminal:
```
make server
```

---
