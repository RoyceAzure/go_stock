package logservice

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	repository "github.com/RoyceAzure/go-stockinfo-logger/repository/mongodb"
	"github.com/rs/zerolog"

)

var Logger zerolog.Logger

type MongoLogger struct {
	mongoDao repository.IMongoDao
}

func SetUpMutiMongoLogger(mongoLogger *MongoLogger) error {
	if mongoLogger == nil {
		return fmt.Errorf("mongo logger is not init")
	}
	multiLogger := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, mongoLogger)
	Logger = zerolog.New(multiLogger).With().Timestamp().Logger()
	return nil
}

func NewMongoLogger(mongoDao repository.IMongoDao) *MongoLogger {
	return &MongoLogger{
		mongoDao: mongoDao,
	}
}

func (mw *MongoLogger) Write(p []byte) (n int, err error) {
	// Insert the record into the collection.

	if mw == nil {
		return 0, fmt.Errorf("mongo logger is not init")
	}

	var logentry repository.LogEntry

	err = json.Unmarshal(p, &logentry)
	if err != nil {
		return 0, err
	}
	err = mw.mongoDao.Insert(context.Background(), logentry)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
