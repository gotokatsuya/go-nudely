# Nudely
Nudity Detection

```
go get github.com/gotokatsuya/go-nudely/cmd/nudely
```

## How to use it

```go
var src image.Image
if src = nudely.DecodeImageByPath(path); src == nil {
	return
}

// Can read file
// if src = nudely.DecodeImageByFile(file); src == nil {
//   return
// }

// detected is bool
detected := nudely.Detect(src)
```

## Execution
```
nudely -path="test.jpg"
```

### Result

When nudely could detect, print
```
Rating : 0.547356
I think this is nude.
```
When nudely could not detect, print
```
Rating : 0.000013
No nude.
```
