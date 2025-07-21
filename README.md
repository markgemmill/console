### console 

A very basic terminal printer that allows for inline colorization
and output levels.

```go
import "github.com/markgemmill/console"


var csl = console.NewConsole(console.DEBUG)

csl.Print("basic message")
csl.Info("message [yellow]with COLOR[/yellow]!")
csl.Debug("[red]DEBUG[/red] %s\n", "message")
csl.Trace("this wont print")

csl.SetLevel(console.TRACE)
csl.Trace("unless the level is set to trace.")

```
```
```
