
# A minimal REST API using Go and MySQL.

## Features
 - get product/category using `GET`
 - list products/categories using `GET`
 - create product/category with `POST`
 - update product/category using `PATCH`
 - listing supports pagination with page/per_page and order by asc/desc 
 - `POST`, `DELETE` and `UPDATE` require authorization. Just use [username and pass](https://github.com/nmakro/best-price-api/blob/master/server/server.go#L18) from basicAuth used inside the server initialization.


## Install

### Get the app

 - $ `go get github.com/nmakro/best-price-api`

Or if you prefer to git clone and manually install the project dependencies:

 - $ `git clone https://github.com/nmakro/best-price-api.git`
 
### Setup the DB
Default configuration expects a running instance of mySQL running on port `3306`. You can change that to match your setup.

Import the example dump file or create a new db named `best_price` with a user name best-price and password `123`.

 
## Use

### Basic usage

$ `go run main.go`

#### Fetch Single Product
$ `curl -i "http://localhost:3000/best-price-api/v1/products/25"`
              
HTTP/1.1 200 OK
Date: Thu, 30 Apr 2020 10:20:06 GMT
Content-Type: application/json
Content-Length: 191

```json
{
    "id": 25,
    "category_id": 5,
    "title": "title 24",
    "image_url": "",
    "price": 72,
    "description": "description 24",
    "deleted_at": null,
    "created_at": "2020-04-30T10:16:34Z",
    "updated_at": "2020-04-30T10:16:34Z"
}
```

### List products from Page 1 with 4 results with Ordering

$ `curl -i "http://localhost:3000/best-price-api/v1/products?per_page=4&page=1&order=price:desc"`
HTTP/1.1 200 OK
Date: Thu, 30 Apr 2020 10:29:11 GMT
Content-Type: application/json
Content-Length: 896

```json
{
    "_meta": {
        "total_records": 400,
        "total_pages": 100,
        "per_page": 4,
        "page": 1,
        "prev_page": 1,
        "next_page": 2
    },
    "products": [{
        "id": 400,
        "category_id": 5,
        "title": "title 199",
        "image_url": "",
        "price": 597,
        "description": "description 199",
        "deleted_at": null,
        "created_at": "2020-04-30T10:18:14Z",
        "updated_at": "2020-04-30T10:18:14Z"
    }, {
        "id": 399,
        "category_id": 4,
        "title": "title 198",
        "image_url": "",
        "price": 594,
        "description": "description 198",
        "deleted_at": null,
        "created_at": "2020-04-30T10:18:14Z",
        "updated_at": "2020-04-30T10:18:14Z"
    }, {
        "id": 398,
        "category_id": 3,
        "title": "title 197",
        "image_url": "",
        "price": 591,
        "description": "description 197",
        "deleted_at": null,
        "created_at": "2020-04-30T10:18:14Z",
        "updated_at": "2020-04-30T10:18:14Z"
    }, {
        "id": 397,
        "category_id": 2,
        "title": "title 196",
        "image_url": "",
        "price": 588,
        "description": "description 196",
        "deleted_at": null,
        "created_at": "2020-04-30T10:18:14Z",
        "updated_at": "2020-04-30T10:18:14Z"
    }]
}
```
