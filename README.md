# Варианты запуска

Вопрос - почему не работает в докере

## Как есть

```bash
go run ./hw/cmd/main.go --config ./hw/configs/http.yaml 

   find me on address - localhost:8080
   Http was running...

```

Проверка

```bash
curl localhost:8080
   
   I receive teapot-status code!

```

## Компиляция

```bash
CGO_ENABLED=0 go build -ldflags "-s -w" -o ./bin/http ./hw/cmd/main.go
./bin/http --config ./hw/configs/http.yaml 

   find me on address - localhost:8080
   Http was running...

```

Проверка

```bash
curl localhost:8080
   
   I receive teapot-status code!

```

## Dockerfile - почему не работает

```bash
docker build --no-cache -t hw/http:v1 .
docker images

   REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
   hw/http      v1        11776bad593a   6 minutes ago   11.6MB
   <none>       <none>    2b06aad0f55b   6 minutes ago   1.74GB
   <none>       <none>    49b66a16e8fb   9 minutes ago   11.6MB
   <none>       <none>    299e9407585f   9 minutes ago   1.74GB
   golang       1.19      80b76a6c918c   11 hours ago    1.06GB
   alpine       3.9       78a2ce922f86   3 years ago     5.55MB

```

Без проброса портов

```bash

docker run -d --name hwhttp hw/http:v1 .

   cae9611478729b249d36ba6f41c2396cebada5cdf71a87ffe747506363310926

docker ps

   CONTAINER ID   IMAGE        COMMAND                  CREATED          STATUS          PORTS     NAMES
   cae961147872   hw/http:v1   "/opt/service/servic…"   25 seconds ago   Up 24 seconds             hwhttp

docker logs hwhttp

   find me on address - localhost:8080
   Http was running...
   
```

Проверка

```bash
curl localhost:8080

   curl: (7) Failed to connect to localhost port 8080 after 0 ms: Couldn't connect to server

```

С пробросом портов

```bash
docker rm -f hwhttp

docker run -d -p 8080:8080 --name hwhttp hw/http:v1 .
   
   bfa2e296cc37190982a15cda25a0261bd97a218583da426db1cd748d61a299db

docker ps

   CONTAINER ID   IMAGE        COMMAND                  CREATED          STATUS          PORTS                                       NAMES
   6145fd9bed13   hw/http:v1   "/opt/service/servic…"   24 seconds ago   Up 23 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   hwhttp

docker logs hwhttp

   find me on address - localhost:8080
   Http was running...

```

Проверка

```bash
curl localhost:8080
   
   curl: (56) Recv failure: Соединение разорвано другой стороной

```
