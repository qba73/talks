# Negating assertion with "!"

`cat.txtar`

```
! exec cat not_existing_file.txt
stderr 'cat: not_existing_file.txt: No such file or directory'
```

Run test:

```bash
go test -run TestScript/cat
PASS
ok  demo 0.154s
```
