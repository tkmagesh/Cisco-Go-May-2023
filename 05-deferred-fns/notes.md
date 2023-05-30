# Deferred Functions #
- postpone the execution of a function until the current function execution is completed
- the deffered functions are executed **irrespective** of how you return from the function

```
    func processData(){
        open the file
       <!--  defer func(){
            close the file
        }() -->
        while not eof {
            read a line
            if err != nil {
                close the file // => logic that we want to execute irrespective of how we return function function
                return
            }
            parse the line
            if err != nil {
                close the file // => logic that we want to execute irrespective of how we return function function
                return
            }
            process the line
            if err != nil {
                close the file // => logic that we want to execute irrespective of how we return function function
                return
            }
        }
    
        close the file
        return result
    }
```