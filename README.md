# loyo-twirp-lab


An experiment to use Twirp and DuckDB for an imaginary service for loyalty programs.


## Building and running

```
$ git clone https://github.com/zikani03/loyo-twirp-lab

$ cd loyo-twirp-lab

$ cd cmd/server

$ LIB="$(pwd)/libduckdb.so" CGO_LDFLAGS="-L$(pwd)" LD_LIBRARY_PATH=$(pwd) CGO_CFLAGS="-I$(pwd)" go build

$ export LD_LIBRARY_PATH=$(pwd)

$ ./server

```


