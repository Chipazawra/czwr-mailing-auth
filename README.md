# Сервис AUTH
____

Мини сервер авторизации. После успешной базовой авторизации на ендпоинте **host:port/auth/login** выдаёт acces и refresh JWT-токены.
На ендпоинте host:port/logout аннулирует токены.

Для сборки и запуска используй  **make**:
- **make build_and_run** инициализирует swagger, собирает проект и запускает на host:port по умолчанию. **127.0.0.1:8885**
- **make build_docker** собирает docker образ. 

Для конфигурирования сервера и спользуются премененные окружения:
- **SIGNING_KEY** - ключ для подписывания  JWT-токенов, также для установки можно использовать флаг **--port**
- **AUTH_HOST** - сервер на котором небходимо запустить сервис, также для установки можно использовать флаг **--host**
- **AUTH_PORT** - порт на котором небходимо запустить сервис   

В сервис добавлен SWAGGER **host:port/auth/swagger/index.html** и **pprof** с возможность вклюичь и отключить профилирование без остановки сервиса с помошью POST запроса **host:port/auth/pprof_enable** и **host:port/auth/pprof_disable**. 
____

Чтобы избежать хранинение паролей в бд сделал обёртку для gin.BasicAuth которая позволяет хранить пароли в md5
____

Настроен небольшой CI при пуше в мастер: lint->test->build Docker-image->[push image on Docker-hub](https://hub.docker.com/repository/docker/chipazawra/czwr-mailing-auth)