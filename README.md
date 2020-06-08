# API Test
API with 2 endpoints:
You can use Postman to test, there is `api_test.postman_collection.json` with the collection of calls.

## Add some Item to a Cart
Add a item to a cart, if the cart doesn't exist it will be created.
```
POST http://localhost:8000/v1/cart/item

```
Json Body:
```json
{
	"cart_id": "1751c582-b8cf-44a2-80ae-735d486b7382",
	"id": "f6e2a013-1b83-4ccd-be41-69ed71a02ce8",
	"name": "Item test 3",
	"description": "The first item added to the cart.",
	"price": 120.5
}
```

## Get Cart Information
Get all the Items of a Cart.
```
GET http://localhost:8000/v1/cart/{cart_id}

```
Json Response body example:
```json
{
    "id": "1751c582-b8cf-44a2-80ae-735d486b7382",
    "items": [
        {
            "id": "f6e2a013-1b83-4ccd-be41-69ed71a02ce8",
            "cart_id": "1751c582-b8cf-44a2-80ae-735d486b7382",
            "name": "Item test 3",
            "description": "The first item added to the cart.",
            "price": 120.5,
            "created_at": "2020-06-08T18:36:06.8439433Z"
        }
    ],
    "created_at": "2020-06-08T18:36:06.8439271Z"
}
```


## About the code
This project has been created using Golang without any aspiration, only for practice.
As repository the API uses the memory, the data won't be persisted. All data will be lost when the Api closes.
I used `dingo` as Dependency Injection solution.  

### Note
In order to make modifications you need to install 
```
go get -u github.com/elliotchance/dingo
```
This is the DI used in the project and it generates the file `dingo.go` using `dingo.yml` as configuration.
To update the file `dingo.go` you must run `./dingo` command in the terminal.

## Get Up the Api
```
docker-compose up --build
```

## Run test
Execute the next command:
``` 
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```
