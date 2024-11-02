# svyatoslave

## instructions

* First of all add node_modules by command:

```
cd client
npm i
```

* Run client app by command:

```
npm start
```

* Run server and try to send some requests
```
cd server
go run main.go
```

## api

| URL-адрес | HTTP-метод | Описание | Тело (Контент) |
|:----------------|:---------:|:----------------|:-----------------:|
|/|GET|Получение сообщения о ближайшем будущем хохлов||
|api/data/|GET|Тестовая ручка для проверки авторизации|Заголовок X-API-KEY(содержимое произвольное)|