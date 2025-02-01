# Comparing output with golden files

`hello.txtar`

```
# With an argument, greet a Gopher and compare with file content.
exec demo Kate
cmp stdout golden.txt

-- golden.txt --
Hello Gopher, Kate!
```

Run test:

```bash
$ go test -run TestHello/hello
PASS
ok   demo 0.115s
```

Run test in verbose mode:

```bash
$ go test -run TestHello/hello -v
=== RUN   TestHello
=== RUN   TestHello/hello
=== PAUSE TestHello/hello
=== CONT  TestHello/hello
  ...

        # With an argument, greet a Gopher and compare with file content. (0.005s)
        > exec demo Kate
        [stdout]
        Hello Gopher, Kate!
        > cmp stdout golden.txt
        PASS

--- PASS: TestHello (0.00s)
    --- PASS: TestHello/hello (0.01s)
PASS
ok   demo 0.114s
```
