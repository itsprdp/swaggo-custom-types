# swaggo-custom-types
Reproduce swaggo issue https://github.com/swaggo/swag/issues/1077


### Run the app on http://localhost:8000
```bash
go run main.go
```

### swaggo generate files
```bash
swag init

2022/12/23 14:53:22 Generate swagger docs....
2022/12/23 14:53:22 Generate general API Info, search dir:./
2022/12/23 14:53:22 Generating main.sampleResponse
2022/12/23 14:53:22 Generating main.CustomType
2022/12/23 14:53:22 ParseComment error in file /Users/itsprdp/dev/swag-bug/main.go :main.subType is not basic types
```
