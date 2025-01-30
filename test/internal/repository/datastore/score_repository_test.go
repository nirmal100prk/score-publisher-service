package datastore

import (
	"context"
	"errors"
	"score-publisher-svc/internal/repository/datastore"
	postgres_mock "score-publisher-svc/test/internal/repository/postgres"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertScore_Success(t *testing.T) {
	// Create a new mock repository
	mockRepo := new(postgres_mock.MockPgxRepository)

	mockTag := pgconn.CommandTag(pgconn.NewCommandTag("INSERT 1"))
	mockRepo.On("Exec", mock.Anything, "INSERT INTO scores (score)  VALUES (  $1)", mock.Anything).Return(mockTag, nil)

	// Initialize the data repository with the mock
	dataRepo := datastore.NewDataRepository(mockRepo)

	// Call the actual InsertScore func
	err := dataRepo.InsertScore(context.Background(), 100)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestInsertScore_Failure(t *testing.T) {
	mockRepo := new(postgres_mock.MockPgxRepository)

	mockRepo.On("Exec", mock.Anything, "INSERT INTO scores (score)  VALUES (  $1)", mock.Anything).
		Return(pgconn.CommandTag(pgconn.NewCommandTag("")), errors.New("database error"))

	dataRepo := datastore.NewDataRepository(mockRepo)

	err := dataRepo.InsertScore(context.Background(), 100)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestInsertScore_NoRowsAffected(t *testing.T) {

	mockRepo := new(postgres_mock.MockPgxRepository)

	mockRepo.On("Exec", mock.Anything, "INSERT INTO scores (score)  VALUES (  $1)", mock.Anything).
		Return(pgconn.CommandTag(pgconn.NewCommandTag("")), nil)

	dataRepo := datastore.NewDataRepository(mockRepo)

	err := dataRepo.InsertScore(context.Background(), 100)

	assert.Error(t, err)
	assert.Equal(t, "query failed 0 rows affected", err.Error())

	mockRepo.AssertExpectations(t)
}
