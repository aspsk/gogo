# Exercise 8.1

Modify `clock2` to accept a port number, and write a program, `clockwall`, that
acts as a client of several clock servers at once, reading the times from each
one and displaying results in a table. To test you can run several instances
with fake time zones:
```
TZ=US/Eastern     ./clock2 --port 8881
TZ=Asia/Tokio     ./clock2 --port 8882
TZ=Europe/London  ./clock2 --port 8883
```
