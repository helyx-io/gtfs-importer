package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQL
////////////////////////////////////////////////////////////////////////////////////////////////

func openSqlConnection() (gorm.DB, error) {
	return gorm.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8")
}


func InitDb(maxIdelConns, maxOpenConns int) (*gorm.DB, error) {
	db, err := openSqlConnection()

	if err != nil {
		return nil, err
	}

	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()

	db.DB().SetMaxIdleConns(maxIdelConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.SingularTable(true)

	return &db, nil
}


func CreateMySQLGTFSRepository(db *gorm.DB) database.GTFSRepository {
	return MySQLGTFSRepository{db}
}

type MySQLGTFSRepository struct {
	db *gorm.DB
}

type MySQLGTFSModelRepository struct {
	db *gorm.DB
}

type MySQLImportTask struct {
	tasks.ImportTask
	db *gorm.DB
}
