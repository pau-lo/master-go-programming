## Naming conventions

Naming conventions are important for code readability and maintainability

- Name starts with a letter or an underscore
- cases matter
- Go keywords cannot be used as names
- use the first letters of the words

For example:

`var mv int // mv -> max value`

- use fewer letter in smaller scopes and the complete word in larger scopes

`var packetsReceived int // Is not ok, too verbose`

`var n int // is ok.  number of packets Received`

`var taskDone bool // Is ok in letter scopes`