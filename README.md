# Go Authorization server

To make life easier, you can run the project with [Compile Daemon](https://github.com/githubnemo/CompileDaemon).

```bash
CompileDaemon --command="./go-auth-server"
```

This project is powered by [GIn](https://gin-gonic.com/docs/quickstart/) and [GORM](https://gorm.io/) with SQLLite.

## Testing

This repository is already included a Postman JSON file, which contains all needed sample requests for every route. You just need to take it from `postman` folder, and import the collection to your Postman application.