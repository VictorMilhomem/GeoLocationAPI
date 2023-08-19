# Backend Challenge

```
{
  "id": 1, 
  "tradingName": "Adega da Cerveja - Pinheiros",
  "ownerName": "Zé da Silva",
  "document": "1432132123891/0001",
  "coverageArea": { 
    "type": "MultiPolygon", 
    "coordinates": [
      [[[30, 20], [45, 40], [10, 40], [30, 20]]], 
      [[[15, 5], [40, 10], [10, 20], [5, 10], [15, 5]]]
    ]
  },
  "address": { 
    "type": "Point",
    "coordinates": [-46.57421, -21.785741]
  }
}
```

1.2. Load partner by id:
Return a specific partner by its id with all the fields presented above.

1.3. Search partner:
Given a specific location (coordinates long and lat), search the nearest partner which the coverage area includes the location.

1.4. Technical Requirements:
The programming language and the database engine are entirely up to you;
Your project must be cross-platform;
Provide a documentation (README.md) file explaining how to execute your service locally and how to deploy it (focus on simplicity, and don't forget that we should test your service on our own, without further assistance).


### Up a docker container

Go to the src folder and run the docker compose command
```
docker-compose up
```

### Run the API

In the src folder run the following command, the server will be in localhost:8000 
```
go run main.go
```