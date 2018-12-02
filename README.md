# serverless-task

`Environment`

```sh
$ mv serverless.env.example.yml serverless.env.yml
```

`Install serverless globally`

```sh
$ npm install serverless -g
```

`Add credentials`

```sh
$ serverless config credentials --provider aws --key AWS_KEY --secret AWS_SECRET
```

`Add dependencies`

```sh
$ npm install
```

`Install goose`

```sh
$ go get -u bitbucket.org/liamstask/goose/cmd/goose
```

`Environment Migrations`

```sh
$ mv db/dbconf.example.yml db/dbconf.yml
```

`Migrations UP`

```sh
$ goose -env development up
```

`Install dep`

```sh
$ go get -v github.com/golang/dep/cmd/dep
```

`Install the project's dependencies`

```sh
$ dep ensure
```

`Build`

```sh
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/api api/main.go
```

`Deploy`

```sh
$ serverless deploy
```