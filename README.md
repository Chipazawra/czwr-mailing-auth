# Сервис AUTH
____

Мини сервер авторизации. После успешной базовой авторизации на ендпоинте host:port/login выдаёт acces и refresh JWT-токены.
На ендпоинте host:port/logout аннулирует токены.

Для сборки и запуска используй  **make**:
- **make build_and_run** инициализирует swagger, собирает проект и запускает на host:port по умолчанию.
- **make build_docker** собирает docker образ. 

Для конфигурирования сервера и спользуются премененные окружения:
- **SIGNING_KEY** - ключ для подписывания  JWT-токенов, также для установки можно использовать флаг **--port**
- **AUTH_HOST** - сервер на котором небходимо запустить сервис, также для установки можно использовать флаг **--host**
- **AUTH_PORT** - порт на котором небходимо запустить сервис   


