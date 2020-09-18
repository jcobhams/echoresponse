## Echo Framework Response Formatter
This is a VERY opinionated package to format HTTP responses.
It absolutely depends on [Echo Framework](github.com/labstack/echo)

### Response Format
It returns JSON responses only
Two keys are a present in the JSON string.

`msg => string`

`payload => interface{}`

`payload` could be anything. Map, List, Int, String


### Why?
I got tired of writing and duplicating the same code across most of my projects.

_Use it if you want/like/don't like :)_