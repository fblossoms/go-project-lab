```sh
curl http://localhost:8080/hello 
```

- 创建书籍
```sh
Invoke-RestMethod -Method POST -Uri "http://localhost:8080/api/books" -ContentType "application/json" -Body '{"title":"The Little Prince","author":"Antoine de Saint-Exupéry","price":29.90,"is_sale":true}'
```

- 获取所有书籍
```sh
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/api/books"
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/api/books?page_number=1&page_size=10"
```

- 根据ID获取书籍
```sh 
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/api/books/1"

```

- 更新书籍
```sh
Invoke-RestMethod -Method PUT -Uri "http://localhost:8080/api/books/1" -ContentType "application/json" -Body '{"title":"The Great Gatsby","author":"F. Scott Fitzgerald","price":39.90,"is_sale":true}'
```

- 删除书籍
```sh
Invoke-RestMethod -Method DELETE -Uri "http://localhost:8080/api/books/1"
```