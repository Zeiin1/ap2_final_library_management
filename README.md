# Api Gateway
Create `dev.env` file in `pkg/config/envs` __(create directory envs)__, then paste the code:

```env
PORT=:3000
AUTH_SVC_URL={svc_host_1}:8080
ORDER_SVC_URL={svc_host_2}:8081
BOOK_SVC_URL={svc_host_3}:8082
// svc_host_1 = svc_host_2 = svc_host_3 = localhost -> in local machine (pc)
// svc_host_1 = auth-svc, 
// svc_host_2 = order-svc, 
// svc_host_3 = book-svc -> in Docker
```

Exchange the values in _{} brackets_ with needed.

---

To run the Api Gateway, write it in terminal:
```
make server
```

---
