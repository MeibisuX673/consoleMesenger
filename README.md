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

