package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/go-redis/redis/v8"
    "context"
    "milo-ia/internal/config"
	"milo-ia/internal/models"
)

var (
    DB   *gorm.DB
    RedisClient *redis.Client
    Ctx = context.Background()
)

func ConnectDatabase(cfg config.Config) error {
    var err error
    dsn := "host=" + cfg.DBHost + " user=gorm password=gorm dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    RedisClient = redis.NewClient(&redis.Options{
        Addr: cfg.RedisAddr,
    })

    return nil
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&models.User{}, &models.Message{})
}
