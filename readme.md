# Задание к модулю

Задача заключается в разработке сервера для сайта новостей. Назовём наш проект GoNews. Наш сайт предоставляет пользователям веб- и мобильное приложение, которое получает данные от сервера по API, и на основе полученных данных отрисовывает пользовательский интерфейс. Данные от сервера приходят в формате JSON.

## ТребованияЖ

1. Сервер GoNews, который мы разрабатываем, должен предоставлять REST API, позволяющий выполнять следующие операции:
    - Получение списка всех статей из БД,
    - Добавление статьи в БД,
    - Обновление статьи в БД,
    - Удаление статьи из БД.

2. Сервер должен предоставлять данные в ответ на запросы по протоколу HTTP.
3. Сервер должен использовать характерную для REST API схему запросов:
    - Запросы должны приходить на URL, соответствующий коллекции ресурсов. Например, коллекция статей(`server/posts`).
    - Для обозначения действий над коллекцией должны использоваться методы протокола HTTP: POST для создания ресурса, DELETE для удаления, PUT для обновления и GET для получения данных.
4. Сервер должен хранить всю информацию в базе данных.
5. Сервер должен предоставить как минимум две реализации хранилища данных: одну для реляционной СУБД и одну для документной.
6. Объекты статьи должны содержать следующую информацию:
    - Идентификатор,
    - Имя автора,
    - Заголовок,
    - Текст,
    - Время создания. 

[ERD](./ERD.png)

# Задание

Для решения задачи требуется следующее:
1. Разработать схему БД PostgreSQL в форме SQL-запроса. Запрос должен быть помещён в файл schema.sql в корневой каталог проекта.
2. По аналогии с пакетом "memdb" разработать пакет "postgres" для поддержки базы данных под управлением СУБД PostgreSQL.
3. По аналогии с пакетом "memdb" разработать пакет "mongo" для поддержки базы данных под управлением СУБД MongoDB.
