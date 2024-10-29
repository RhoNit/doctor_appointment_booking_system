package db

import (
	"context"

	"github.com/RhoNit/doctor_appointment_system/internal/config"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitPgxConn(logger *zap.Logger) (*pgx.Conn, error) {
	dbConfig := config.LoadDBConfig()

	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	&dbConfig.User,
	// 	&dbConfig.Password,
	// 	&dbConfig.Host,
	// 	&dbConfig.Port,
	// 	&dbConfig.Database,
	// )

	connStr := "postgres://postgres:passwd@localhost:5432/echodb?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		logger.Error("Error while connecting to the Database", zap.Error(err))
		return nil, err
	}

	logger.Info("Connected to the database", zap.String("DB", dbConfig.Database))

	return conn, nil
}
