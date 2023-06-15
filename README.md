# Auth Microservice
Create `dev.env` file in `pkg/config/envs` (create directory envs), then put it the code:
```
PORT=:8080
DB_URL=postgres://{username}:{password}@{host}:5432/services
JWT_SECRET_KEY={some_value} (ex. r43t18sc)
```
Exchange the values in {} brackets with neede. 

() BRACKETS ARE EXAMPLE
