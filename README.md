# TinyGo Button Wrapper
Wrapper for adding functionality onto single pin buttons in TinyGo

Built for use with the Adafruit Circuit Playground Express, but should (theoretically) work for any device with a single pin button.

## Example Usage

```go
A := button.New(machine.BUTTONA)
for {
  a, d := A.Held()
  println("a:", a, "| held duration:", d.Milliseconds(), "ms")
  time.Sleep(100 * time.Millisecond)
}
```

## TinyGo extension

Not mandatory, but strongly recommended. Available [here](https://tinygo.org/docs/guides/ide-integration/vscode/).
