# Thinksurance Case Study - Ahmet Can Aydemir

## Libraries

- [gofiber/fiber](https://github.com/gofiber/fiber): Web framework
- [stretchr/testify](https://github.com/stretchr/testify): Mocking
- [spf13/cobra](https://github.com/spf13/cobra):spf13/cobra: Cli commands

## Authentication

You have to authenticate with `Basic Auth` for accessing endpoints. We are using a fake DB for users at [gateway_svc/db/users.go](./gateway_svc/db/users.go).

Username: **admin**\
Password: **admin**

## Run

You can run with docker compose or with minikube.

### Docker Compose

Run following command in terminal and connect gateway via http://127.0.0.1:8080

```bash
docker compose up --build
```

### Minikube

Run following commands in terminal and connect to API gateway via http://127.0.0.1:8080

```
minikube start

minikube image build -t json json_svc/.
kubectl apply -f json_svc/json-deployment.yaml
kubectl apply -f json_svc/json-service.yaml

minikube image build -t algorithm algorithm_svc/.
kubectl apply -f algorithm_svc/algorithm-deployment.yaml
kubectl apply -f algorithm_svc/algorithm-service.yaml

minikube image build -t gateway gateway_svc/.
kubectl apply -f gateway_svc/gateway-deployment.yaml
kubectl apply -f gateway_svc/gateway-service.yaml

kubectl port-forward service/gateway-service 8080:8080
```

## Endpoints

### Algorithm

> POST /find-position

This endpoint returns position of searched integer in sorted array. If not found it returns -1.

#### Request Body

| Field                 | Type  | Description                         |
|-----------------------|-------|-------------------------------------|
| `array`    (required) | []int | Sorted integer array                |
| `search`(required)    | int   | Integer to search position in array |

#### Example Successful Request

```json
POST /find-position
{
    "array": [1,1,2,3,3,4,5,6],
    "search": 3
}
```

#### Example Successful Response

```json
4
``` 

### JSON
> GET /persons

This endpoint reads json files in persons folder, converts to `Person` objects and returns the array.

#### Example Successful Request

```json
GET /persons
```

#### Example Successful Response

```json
[
   {
      "FirstName": "Coy",
      "LastName": "Mertz",
      "BirthDay": "1944-09-14T06:32:39.910Z",
      "Address": "8742 Stehr Trail Apt. 220",
      "PhoneNumber": "559-487-641"
    },
    {
    "FirstName": "Kate",
    "LastName": "Musk",
    "BirthDay": "1944-09-14T06:32:39.910Z",
    "Address": "8742 Stehr Trail Apt. 220",
    "PhoneNumber": "559-487-641"
    },
  ...
]
```

