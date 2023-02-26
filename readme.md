# Simple Project that used for Completing Techical Test as Backend Developer Intern in Dagangan

## Database
<p>This Project is only used 1 table with 8 columns</p>
<p>There are id, name, description, price, quantity, created_at, updated_at, deleted_at</p>

## Route For this Project
1. GET /product/:id => to get product with ID 
2. GET /products => to get all products 
3. POST /product => to create new Product
4. PUT /product/:id => to edit product
5. DELETE /product/:id => to soft delete product (Add value on deleted_at column)

### POST Request
<p>I didn't add validation rule on post request, So when you posted a new data with no value, it will be saved as well</p>
<p>If you want to post data with values, add keys and values to form-data and the keys to use are name, description, price, quantity<p>


### Pagination 
<p>Pagination is only used on /products route</p>
<p>If you want to use pagination, add "limit" and/or "offset" as params</p>

### Filtering
<p>Filtering is only used on /products route</p>
<p>If you want to use filtering, add "filter" as params</p>

#### **If you want to use pagination and filtering together, add "limit", "offset", and "filter"** as params

