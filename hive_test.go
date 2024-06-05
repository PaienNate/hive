package hive

import (
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dbDSN = "gorm:gorm@localhost:10000?auth=PLAIN"

func TestMain(m *testing.M) {
	Init()
	os.Exit(m.Run())
}

func Init() {
	var err error
	if DB, err = gorm.Open(Open(dbDSN), &gorm.Config{}); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}

	RunMigrations()

	DB = DB.Debug()

}

func RunMigrations() {
	allModels := []interface{}{&User{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	if err := DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err := DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}

}

func TestList(t *testing.T) {
	sql := "INSERT INTO `users`" +
		"(`id1`, `name`,`age`, `active`, `salary`," +
		" `created_at`, `updated_at`, `date`, `score`)" +
		"SELECT 1, 'philhuan', 1,true, 100, " +
		"'2024-05-04 12:35:30', '2024-05-04 12:35:30', '2024-05-04', MAP('math',120,'english', 123);"
	ret := DB.Model(&User{}).Exec(sql)
	//assert.Equal(t, 1, ret.RowsAffected) // TODO
	assert.NoError(t, ret.Error)

	var res []User
	err := DB.Table("users").Find(&res).Error
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestNew(t *testing.T) {
	dialector := New(Config{DSN: dbDSN})
	d, ok := dialector.(*Dialector)
	if !ok {
		t.Fatal("dialector is not *Dialector")
	}
	if d.DSNConfig == nil {
		t.Error("dialector.DSNConfig is nil")
	}

	dialector = New(Config{DSNConfig: &DSNConfig{}})
	d, ok = dialector.(*Dialector)
	if !ok {
		t.Fatal("dialector is not *Dialector")
	}
	if d.DSN == "" {
		t.Error("dialector.DSN is empty")
	}
}
