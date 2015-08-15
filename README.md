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

detect := nudely.Detect(src)
```

## Execution
```
nudely -path="test.jpg"
```

### Result

It could detect
```
Rating : 0.547356
I think this is nude.
```
It could not detect
```
Rating : 0.000013
No nude.
```
