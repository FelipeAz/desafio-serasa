package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/FelipeAz/desafio-serasa/internal/pkg/app/entity"
	"github.com/FelipeAz/desafio-serasa/internal/pkg/app/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLHandler contem a conexao do banco de dados.
type SQLHandler struct {
	db *gorm.DB
}

// NewSQLHandler inicia uma conexao com o Banco de dados SQL.
func NewSQLHandler() (interfaces.SQLHandler, error) {
	dbHandler := &SQLHandler{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/Serasa?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbHandler.db = db
	err = dbHandler.migrateTables()

	return dbHandler, err
}

// CloseConnection fecha a conexao com o banco de dados.
func (env *SQLHandler) CloseConnection() {
	sql, err := env.db.DB()
	if err != nil {
		log.Println(err)
	}

	sql.Close()
}

// GetGorm retorna uma instancia do GORM que sera utilizada para busca no BD.
func (env *SQLHandler) GetGorm() *gorm.DB {
	return env.db
}

func (env *SQLHandler) migrateTables() (err error) {
	err = env.db.Migrator().AutoMigrate(
		&entity.Negativacao{},
		&entity.User{},
		&entity.Access{},
	)

	return
}
