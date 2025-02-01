# Comparing output using `regex`

`hello.txtar`

```
# With an argument, greet a Gopher and compare with file content.
exec demo Kate
exists greet.out
cmp greet.out golden.txt

# With an argument, check if the argument exist in the file content.
exec demo Michael
exists greet.out
grep 'Michael' greet.out

-- golden.txt --
Hello Gopher, Kate
```

Run test:

```bash
go test -run TestHello/hello
PASS
ok   demo 0.115s
```

Run test in verbose mode:

```bash
go test -run TestHello/hello -v
=== RUN   TestHello
=== RUN   TestHello/hello
=== PAUSE TestHello/hello
=== CONT  TestHello/hello
  ...

        # With an argument, greet a Gopher and compare with file content. (0.005s)
        > exec demo Kate
        [stdout]
        Hello Gopher, Kate

        > exists greet.out
        > cmp greet.out golden.txt
        # With an argument, check if the argument exist in the file content. (0.005s)
        > exec demo Michael
        [stdout]
        Hello Gopher, Michael

        > exists greet.out
        > grep 'Michael' greet.out
        PASS

--- PASS: TestHello (0.00s)
    --- PASS: TestHello/hello (0.01s)
PASS
ok   demo 0.115s
```
