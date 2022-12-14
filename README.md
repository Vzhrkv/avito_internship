# Тестовое задание на позицию стажера бэкенда

<!-- TOC start-->
## Содержание
1. [Описание задания](#Описание-задания)
2. [Реализация](#Реализация)
3. [Методы](#Методы)
4. [Нерешенные вопросы](#Нерешенные-вопросы)
<!-- TOC end-->

## Описание задания 

Реализация микросервиса для работы с балансом пользователей.
Небходимо было реализовать методы: начисления,
списания средств, перевод средств от пользователя к пользователю, метод получения баланса,
метод списания средств.

## Запуск

В начале нужно создать запустить контейнер с Postgres. 
Пример 
```bash
docker run --name=users_balance -e POSTGRES_PASSWORD="qwerty" -p 5435:5432 -d --rm postgres
```
Запуститься в созданном контейнере следующей командой:
```bash
docker exec -it container_id /bin/bash
```

Внутри контейнера ввести команду:
```bash
psql -U postgres
```

И создать таблицы командами:
```bash
create table users (
    user_id int not null unique,
    balance int
);
```

```bash
create table reservedFunds (
    user_id int references users(user_id) not null ,
    service_id int not null,
    order_id int not null,
    price int
);
```

Запуск приложения:
```bash
make run 
```


## Реализация

- Старался следовать чистой архитектуре настолько, насколько смог.
- Из сторонних библиотек использовались: [gorilla/mux](https://github.com/gorilla/mux), 
[spf13/viper](https://github.com/spf13/viper), 
[sirupsen/logrus](https://github.com/sirupsen/logrus).
- В качестве СУБД использовался Postgres с использованием стандартной библиотеки
[database/sql](https://pkg.go.dev/database/sql).

## Методы

Все методы принимают в качестве запроса JSON с определенными полями.
Ответ выдается в виде JSON.

### Добавление средств на баланс пользователя (Создание пользователя)

```bash
curl --location --request POST 'localhost:4303/user/add/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": user_id,
    "add_funds": add_funds
}'
```
#### Поля
`user_id` - ID пользователя, у которого хотим обновить баланс.

`add_funds` - количество средств, зачисляемых на счет.

#### Ответ 
JSON с полями "Status" и "Msg". 

"Status" - это либо "failed" при неудачной попытке добавления, "accepted" при успехе.

"Msg" - поле с сообщением об операции. При неуспешном добавлении выводится сообщение ошибки. 
При успешном добавлении возвращает сообщение "Added funds to user(user_id=user_id)"

### Получение счета пользователя

```bash
curl --location --request GET 'localhost:4303/user/get/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": user_id
}'
```
#### Поля

`user_id` - ID пользователя, у которого хотим получить значение баланса.

#### Ответ
JSON с полями "user_id" и "balance".

"balance" - поле со значением баланса пользователя.

### Перевод средств другому пользователю.

```bash
curl --location --request POST 'localhost:4303/user/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": user_id,
    "other_user_id": other_user_id,
    "funds_to_send": funds_to_send
}'
```

#### Поля

`user_id` - ID пользователя, с которого переводяться средства.

`other_user_id` - ID пользователя, на счет которого переводяться средства.

`funds_to_send` - количество средств для перевода.

#### Ответ 
JSON c полями "Status" и "Msg".

"Status" - это либо "failed" при неудачной попытке перевода, "accepted" при успехе.
"Msg" - в случае неудачи - сообщение об ошибке, при успехе - "Sended funds".

### Резервирование средств 

```bash
curl --location --request POST 'localhost:4303/user/reserve/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": user_id,
    "service_id": service_id,
    "order_id": order_id,
    "price": price
}'
```
#### Поля

`user_id` - ID пользователя, с баланса которого будут зарезервированы средства.

`service_id` - ID услуги.

`order_id` - ID заказа.

`price` - стоимость или количество средств для резервирования.

#### Ответ
JSON c полями "Status" и "Msg".

"Status" - при успехе "accepted", иначе "failed".

"Msg" - при неудаче выводит сообщение об ошибке, при успехе - "Funds for order(user_id=user_id, service_id=service_id, order_id=order_id, price=price) was reserved".

### Подтверждение заказа пользователя

```bash
curl --location --request POST 'localhost:4303/user/confirm-order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": user_id,
    "service_id": service_id,
    "order_id": order_id,
    "price": price
}'
```

#### Поля 
Поля в этом методе имеют тот же смысл, что и в [Резервирование средств](###Резервирование-средств).

#### Ответ 
JSON c полями "Status" и "Msg".

"Status" - при успехе "accepted", иначе "failed".

"Msg" - при неудаче выводит сообщение об ошибке, при успехе - "Order(user_id=33, service_id=31, order_id=23, price=100) was confirmed".

### Возникшие вопросы

#### Где хранить зарезервированные средства?

Самое первое решение - хранить сведения о заказе в таблице `users`. Но такое решение не удобное в том плане, что не оптимально
происходит работа с данными. Все лежит в куче и попробуй в ней разберись.

Поэтому была создана вторая таблица `reservedfunds` , которая содержит в себе информацию о заказе. При подтверждении списания сведения удаляются из таблицы.
Такая структура удобна тем, что при добавлении/удалении данных нужно обращаться только к нужным сведениям и изменять только их.

#### Нерешенные вопросы
1. Подключение контейнера с приложением к контейнеру с Postgresql.
2. Не придуман способ отличить заказы с одинаковыми параметрами, т.е. при списании двух идентичных заказов у одного и того же пользователя, сведения записываются один раз в отчет, а таких записей должно быть несколько.