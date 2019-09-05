# GO run Person API

## Setup

### Install DynamoDB

```sh
docker pull amazon/dynamodb-local
```

### Start DynamoDB

```sh
docker run -p 8000:8000 amazon/dynamodb-local
```

### Create table

```sh
aws dynamodb create-table --cli-input-json file://schema.json --endpoint-url http://localhost:8000
```

### Build and start server

```
grnc-yaml-bind && go build && ./personapi
```

### Access api

#### 1. Get person

**Request:**

```
GET localhost:8080/person
```

**Response**

```json
[
  {
    "age": 21,
    "gender": 2,
    "uid": "6d97db34-cd3c-4c13-a6d0-a57b04953ec8",
    "name": "John Doe"
  },
  {
    "age": 21,
    "gender": 1,
    "uid": "51b9e79f-f596-4be1-8b3d-565e9eb0fa1c",
    "name": "Hari Bahadur"
  }
]
```

#### 2. Create Person

**Request:**

```
POST localhost:8080/person
Params:
{
  "name": "John Doe",
  "age": 20,
  "gender": 1
}
```

**Response**

```json
{
  "age": 20,
  "gender": 1,
  "uid": "6d97db34-cd3c-4c13-a6d0-a57b04953ec8",
  "name": "John Doe"
}
```

#### 3. Get person

**Request:**

```
PUT localhost:8080/person
Params:
{
  "uid": "6d97db34-cd3c-4c13-a6d0-a57b04953ec8",
  "name": "John Doe",
  "age": 22,
  "gender": 1
}
```

**Response**

```json
{
  "age": 22,
  "gender": 1,
  "uid": "6d97db34-cd3c-4c13-a6d0-a57b04953ec8",
  "name": "John Doe"
}
```
