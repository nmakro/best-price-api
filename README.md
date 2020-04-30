<h1>A minimal rest-api using go and MySql.</h1>


<h5>Installation steps:</h5>

<ul>
<li>git clone https://github.com/nmakro/best-price-api.git</li>
 or
<li>go get github.com/nmakro/best-price-api</li>

Note: If using git clone you must manually install the project dependencies.
</ul>

<h5>Requirements:</h5>
A running instance of mySql running on port: 3306.
You can change that later in the configuration.

<h5>Basic usage</h5>

Import the example dump file or create a new db named `best_price` with a user name best-price and password `123`.

From the root directory execute: go run main.go.

Supported actions are:
<ol>
<li>create product/category with POST</li>
<li>update product/category using PATCH</li>
<li>get product/category using GET</li>
<li>list products/categories using GET</li>
<li>listing supports pagination with page/per_page and order by asc/desc </li>
</ol> 

<h5>Examples</h5>

* **Sample Call:**
curl -i "http://localhost:3000/best-price-api/v1/products/25"
              
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

* **Sample List:**
curl -i "http://localhost:3000/best-price-api/v1/products?per_page=4&page=1&order=price:desc"
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


