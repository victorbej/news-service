# Новостной сервис

1. Установите Protobuf
- Руководство для установки на Windows:
  1) Скачать Protocol Buffer для своей операционной системы.
  2) Распаковать архив и переместить файл с расширением .bin на диск C
  3) Скопировать путь и задать его в переменной окружения (Система -> Изменение системных переменных среды -> Переменные среды... -> Два раза кликнуть на PATH -> Создать -> Вставить скопированный путь -> ОК)
  4) Установить go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  5) Установить go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  6) Скачать go get -u github.com/golang/protobuf/protoc-gen-go
  7) Проверить в терминале командой protoc, должен отобразиться текст с информацией о прото.

2. Установите gRPC (Guide - https://grpc.io/docs/languages/go/quickstart/)
2. Склонируйте google api https://fuchsia.googlesource.com/third_party/googleapis/
3. Выполните команду "protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pkg/api/news.proto
   " from root project directory
4. Второй способ генерации protobuf -> buf.build (vpn only)
5. Образец запроса: http://localhost:8081/eapi/news/manage/health -> result is {
   "serviceName": "CheckHealthServer",
   "serviceStatus": "200"
   }