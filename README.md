### Запуск

```bash 
go run tcp-server.go
```

```bash 
go run tcp-client.go [address]
```

### Пример отправки сообщения

[id клиента] сообщение

### Настройки

```-color``` - установить цвет текста <br>

Цвета:
* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white

Пример 
``` bash
go run tcp-client.go -color=green/-color green [address] 
```

Установленный цвет сохраняется

### Сборка исполняемого файла под os

Запуск скрипта 
``` bash
./build/build.sh
```
Чтобы просмотреть список поддерживаемых платформ запустите комманду
``` bash
go tool dist list
```

### Запуск elk

 1. Создайте папку docker_volumes
 2. Укажите в docker/config/filebeat/filebeat.docker.yml имя контейнера с которого хотите собирать логи
 3. Указать в docker/config/filebeat/filebeat.docker ip адресс интерфейса eno1 или eno0
 4. Укажите в env пользователя и пароль, также нужно продублировать в следующих конфигах: 
    - docker/config/kibana
    - docker/config/pipelines/service_stamped_fileBeat_logs
 5. Выполнить комманду docker compose up -d
 