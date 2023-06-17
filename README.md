# Book Microservice
Create `dev.env` file in `pkg/config/envs` __(create directory envs)__, then paste the code:

```env
PORT=:8082
DB_URL=postgres://{username}:{password}@{host}:5432/services
// host = localhost -> in local machine (pc)
// host = postgresql -> in Docker
```

Exchange the values in _{} brackets_ with needed.

---

It will not create a database `services` by itself. I ask you to create it by own!

---

To run the Book Microservice, write it in terminal:
```
make server
```

---
