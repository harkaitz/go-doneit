GO-DONEIT
=========

Execute a function `prog` and if it succeeds, store a record in the database
to avoid executing it again.

The table is `task` with `ID`, `Done` and `DoneDate` fields. Create
it with `InitDatabase`.

## Go programs

    Usage: run-only-once ID COMMAND
    
    Execute command only once, the database is in "~/.run-only-once.db".

## Go documentation

    package doneit // import "github.com/harkaitz/go-doneit"
    
    func InitDatabase(gdb *gorm.DB) (err error)
    func OnlyOnce(gdb *gorm.DB, prog func() error, format string, a ...any) (err error)
    type Task struct{ ... }

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
