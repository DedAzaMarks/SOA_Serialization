Ссылка на Docker-образ https://hub.docker.com/r/dedazamarks/formats-comparison

`docker pull dedazamarks/formats-comparison`

Зависимости
- Go 1.17

Запуск  
`go test -bench=.`

| Format   | Serialize size (B) | Serialize time (ns) | Deserialize size | Deserialize time |
| ---      | ---                | ---                 | ---              | ---              |
| GOB      | 5500               | 56021               | 5521             | 24099            |
| XML      | 12219              | 188630              | 12057            | 612509           |
| JSON     | 7794               | 85074               | 7910             | 146791           |
| PROTOBUF | 5620               | 61911               | 5558             | 30208            |
| AVRO     | 5447               | 60413               | 5394             | 72577            |
| MSGPACK  | 7818               | 100338              | 7821             | 149470           |
| YAML     | 8401               | 676295              | 8456             | 518786           |