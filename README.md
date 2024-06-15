<h1>Запуск:</h1>

<h3>Сервер:</h3>
```bash
sudo go run cmd/server/server.go
```

<h3>Клиент:</h3>
<p>Задать имя хоста:</p>
```bash
go run cmd/client/client.go set-hostname <hostname>
```
<p>Получить список DNS серверов:</p>
```bash
go run cmd/client/client.go list-dns
```
<p>Добавить DNS:</p>
```bash
go run cmd/client/client.go add-dns <8.8.8.8>
```
<p>Удалить DNS:</p>
```bash
go run cmd/client/client.go remove-dns <8.8.8.8>
```
