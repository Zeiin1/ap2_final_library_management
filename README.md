# Auth Microservice
Create `dev.env` file in `pkg/config/envs` __(create directory envs)__, then paste the code:

```env
PORT=:8080
DB_URL=postgres://{username}:{password}@{host}:5432/services
JWT_SECRET_KEY={some_value} (ex. r43t18sc)
// host = localhost -> on local machine (pc)
// host = {docker_db_container_name} -> on Docker
```

Exchange the values in _{} brackets_ with needed.

---

It will not create a database `services` by itself. I ask you to create it by own!

---

To run the Auth Microservice, write it in terminal:
```
make server
```

---
