package postgres_mock

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/mock"
)

// MockPgxRepository is a mock implementation of PgxRepository.
type MockPgxRepository struct {
	mock.Mock
}

func (m *MockPgxRepository) DB() *pgxpool.Pool {
	args := m.Called()
	return args.Get(0).(*pgxpool.Pool)
}

func (m *MockPgxRepository) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	mockArgs := m.Called(ctx, query, args)
	return mockArgs.Get(0).(pgx.Rows), mockArgs.Error(1)
}

func (m *MockPgxRepository) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	mockArgs := m.Called(ctx, query, args)
	return mockArgs.Get(0).(pgx.Row)
}

func (m *MockPgxRepository) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	mockArgs := m.Called(ctx, query, args)
	return mockArgs.Get(0).(pgconn.CommandTag), mockArgs.Error(1)
}

func (m *MockPgxRepository) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	mockArgs := m.Called(ctx, tableName, columnNames, rowSrc)
	return mockArgs.Get(0).(int64), mockArgs.Error(1)
}

func (m *MockPgxRepository) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	mockArgs := m.Called(ctx, b)
	return mockArgs.Get(0).(pgx.BatchResults)
}

func (m *MockPgxRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	mockArgs := m.Called(ctx)
	return mockArgs.Get(0).(pgx.Tx), mockArgs.Error(1)
}

func (m *MockPgxRepository) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	mockArgs := m.Called(ctx, txOptions)
	return mockArgs.Get(0).(pgx.Tx), mockArgs.Error(1)
}

func (m *MockPgxRepository) Close() {
	m.Called()
}
