DUCKDB_LIB_DIR="~/Projects/oss/duckdb/target/"
.phony: echo "Haha"


build:
	CGO_LDFLAGS="-L${DUCKDB_LIB_DIR}/libduckdb_static.a" CGO_CFLAGS="-I/${DUCKDB_LIB_DIR}/duckdb.h" go run examples/simple.go
