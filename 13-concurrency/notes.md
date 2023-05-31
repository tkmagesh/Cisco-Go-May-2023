# Concurrency #

![image Concurrency](./concurrency-model.png)

## Language Features
- go keyword
- channel data type
- channel operator (<-)
- "range" construct
- "select case" construct

## Api Support
- sync package
- sync/atomic package

## Notes
- The scheduled function will be picked up for execution ONLY when the current function execution is blocked for any reason

## To detect race conditions
- go run --race <program.go>
- go build --race <program.go>


