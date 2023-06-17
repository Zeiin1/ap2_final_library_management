# INFO
There is no *database service* in `project/docker-compose.yml`! \
\
Should add it by own.

```yml
// in `project/docker-compose.yml`
// service_name = db_host (be aware)
{service_name}:
  image: {docker_psql_image} (ex. postgres:16beta1)
  ports:
    - "5432:5432"
  environment:
    POSTGRES_USER: {db_name}
    POSTGRES_PASSWORD: {db_password}
    POSTGRES_DB: services
```
