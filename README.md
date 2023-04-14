# SIMPLE IMAGE RESIZING API

Web API to resive images by passing image url, desired width and height. Returns the resized image.

## TOOLS

-   GO
-   Gin (Web server)

## Dependencies

```

Download latest packages

go mod tidy

Try on localhost

go run main.go
```

## How it works

`images?url=[IMAGE_URL]&w=[IMAGE_WIDTH]&h=[IMAGE_HEIGHT]`

`images?url=http://website.com/image.jpg&w=1100&h=1100`

## LIVE DEMO

`(https://imgrsv.herokuapp.com/)`

`(https://imgrsv.herokuapp.com/images?url=https://www.naijaloaded.com.ng/wp-content/uploads/2023/01/wizkidd.jpg&w=400&h=400)`
