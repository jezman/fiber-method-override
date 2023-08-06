# Overrider

Fiber middleware. Overrides POST method to PUT, PATCH od DELETE.

### Usage:

```sh
go get github.com/jezman/overrider
```

### add middleware

```go

import "github.com/jezman/overrider"

func main() {
    app = fiber.New()
    app.Use(overrider.New())
}
```

### add hidden field to your post form
```html
 <input type=hidden name="_method" value="PUT">
 ```
