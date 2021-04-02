# cartesian
Api to Calculate [Manhattan distance](https://xlinux.nist.gov/dads/HTML/manhattanDistance.html#:~:text=Definition%3A%20The%20distance%20between%20two,y1%20%2D%20y2%7C.) in a simple 2-dimensional plane using [Go](https://golang.org/).
This api receive 3 required params:
- `x` integer. This represents the `x` coordinate of the search origin.
- `y` integer. This represents the `y` coordinate of the search origin.
- `distance` integer and positive. This represents the Manhattan distance; points within `distance` from `x` and `y` are returned, points outside are filtered out.


## Requirements

To Run this project you need [Docker](https://www.docker.com/).

## Run

To up the server
```
make run
```
So:

```
curl http://localhost:80/api/points/{x}/{y}/{distance} 
```
example:

```
curl http://localhost:80/api/points/-2/-8/84 
```

## Testing

### Unit tests

Run:

```
make check
```

### Integration tests

Run:

```
make check-integration
```

### Coverage

Run:

```
make coverage
```