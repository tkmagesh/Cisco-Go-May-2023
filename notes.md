# Magesh Kuppan #

# Schedule
    - Commence      : 9:00 AM
    - Tea Break     : 11:00 AM (15 mins)
    - Lunch Break   : 1:00 PM (45 mins)
    - Tea Break     : 3:30 PM (15 mins)
    - Wind up       : 5:00 PM

# Methodology
    - No powerpoints
    - 100% Code driven class
    - Q&A at all times

# Repository
    - https://github.com/tkmagesh/cisco-go-may-2023

# Software Requirements
    - Go Tools (https://go.dev/dl)
    - Visual Studio Code (https://code.visualstudio.com)
    - Go extensions for VS Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)


# Why Golang?
    - Simplicity
        - The paradox of choice (Barry Shwartz)
        - 25 keywords
        - No classes (only structs)
        - No inheritance (only composition)
        - No reference types (everything is a value)
        - No exceptions (only errors)
        - No "try catch finally" construct
        - No access modifiers
        - No pointer arithmatic
        - No implicity type conversion
        - No function overloading
        - No operator overloading
    - Close to hardware
        - Go applications are compiled to machine code
        - Smaller footprint on resources at runtime
    - Concurrency Model
        - Concurrent operations are represented as goroutines (cheap - ~4kb/goroutine)
        - N:M scheduler (can schedule N goroutines using M OS-threads)
        - Concurrency support is built in the language
            - go keyword, channel data type, channel operator (<-), range construct, select-case constuct
        - APIs
            - sync package
            - sync/atomic package
