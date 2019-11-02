## At Once
At Once sends all the request at once.

```bash
$ ./at-once -r 100 -w 5 https://your.server/path
```

| # | option | type |note |
| --- | --- | --- | --- |
| 1 | r | int | total request |
| 2 | w | int | num of workers |
| 3 | m | string | http method (ex. GET / POST / PUT / DELETE / ...) |
| 4 | h | string | http header (ex. "key1:value1,key2:value2" ) |
| 5 | vr | bool | verbose output of response |
| 6 | vt | bool | verbose output of time |


## Per Sec
Per Sec sends specified amount of requests in each second.

```bash
$ ./per-sec -r 20 -s 10 -w 5 https://your.server/path
```

| # | option | type |note |
| --- | --- | --- | --- |
| 1 | r | int | request per sec |
| 2 | s | int | sec |
| 3 | w | int | num of workers (0 means thread per message mode) |
| 4 | m | string | http method (ex. GET / POST / PUT / DELETE / ...) |
| 5 | h | string | http header (ex. "key1:value1,key2:value2" ) |
| 6 | vr | bool | verbose output of response |
| 7 | vt | bool | verbose output of time |
