package loyo

// Loyo is a platform for loyalty
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
	"net/http"
)

type LoyoDataIngestor struct {
	ctx context.Context
	db  *sql.DB
}

func NewIngestor(dbPath string) (*LoyoDataIngestor, error) {
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingestor: %v", err)
	}

	return &LoyoDataIngestor{
		db:  db,
		ctx: context.Background(),
	}, nil
}

func (ingestor *LoyoDataIngestor) CreateTables() error {
	_, err := ingestor.db.ExecContext(ingestor.ctx, `
    CREATE OR REPLACE TABLE loyo_events(
		EventType TEXT,
		TransactionId TEXT,
		ProductSku TEXT,
		ServiceId TEXT,
		AccountId TEXT,
		Currency TEXT,
		Amount DECIMAL(19, 6),
		TaxApplicable DECIMAL(19, 6),
		TaxCharged DECIMAL(19, 6),
		RecipientAccountId TEXT,
		Flagged BOOL,
		Timestamp TEXT,
		ServiceProviderId TEXT
	);
    `)
	if err != nil {
		return err
	}

	return nil
}

func (ingestor *LoyoDataIngestor) Ingest(event *LoyoEvent) error {
	_, err := ingestor.db.QueryContext(context.Background(), `INSERT INTO loyo_events(
		EventType,
		TransactionId,
		ProductSku,
		ServiceId,
		AccountId,
		Currency,
		Amount,
		TaxApplicable,
		TaxCharged,
		RecipientAccountId,
		Flagged,
		Timestamp,
		ServiceProviderId
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		event.EventType,
		event.TransactionId,
		event.ProductSku,
		event.ServiceId,
		event.AccountId,
		event.Currency,
		event.Amount,
		event.TaxApplicable,
		event.TaxCharged,
		event.RecipientAccountId,
		event.Flagged,
		event.Timestamp,
		event.ServiceProviderId,
	)
	if err != nil {
		return err
	}
	return nil
}

type Server struct {
	ingestor *LoyoDataIngestor
}

func (s *Server) IngestEvents(ctx context.Context, req *LoyoIngestRequest) (*LoyoIngestResponse, error) {
	res := &LoyoIngestResponse{}
	for _, event := range req.Events {
		fmt.Println("Attempting to ingest event id:", event.TransactionId)
		err := s.ingestor.Ingest(event)
		if err != nil {
			fmt.Printf("Failed to ingest event, got error %v\n", err)
			return nil, err
		}
	}
	return res, nil
}

func NewServer(bindAddress string) error {
	ingestor, err := NewIngestor("./data/loyo.db")
	if err != nil {
		return err
	}
	err = ingestor.CreateTables()
	if err != nil {
		return err
	}
	server := &Server{
		ingestor: ingestor,
	}

	twirpHandler := NewLoyoIngestServiceServer(server)
	fmt.Println("Server starting ...")
	return http.ListenAndServe(bindAddress, twirpHandler)
}
