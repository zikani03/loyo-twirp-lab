package loyo

// Loyo is a platform for loyalty 
import (
	"fmt"
	"sql"
	_ "github.com/marcboeker/go-duckdb"
)

type LoyoDataIngestor {
	db *sql.DB
}

func NewIngestor(dbPath string) (*LoyoDataIngestor, error) {
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingestor: %v", err)
	}

	return &LoyoDataIngestor { db: db }, nil
}


func (ingestor *LoyoDataIngestor) IngestBatch(events []*pb.LoyoEvent) (int, int, error) {
	result, err := db.ExecContext(ctx, "INSERT INTO loyo_events() VALUES ($1, $2)")

	return 0, 0, fmt.Errorf("not implemented")
}
