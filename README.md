# Nudely
Nudity Detection

```
go get github.com/gotokatsuya/go-nudely/cmd/nudely
```

## How to use it

```go
// img is image.Image, err is error
img, err := nudely.DecodeImageByPath(path)

// Can read file
// img, err := nudely.DecodeImageByFile(file)

// detected is bool, rating is flaot32
detected, rating := nudely.Detect(img)

// If you do not want to know rating
// detected, _ := nudely.Detect(img)
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
Nudely think this is nude when rating is bigger than 0.5 value.  

When nudely could not detect, print
```
Rating : 0.000013
No nude.
```
