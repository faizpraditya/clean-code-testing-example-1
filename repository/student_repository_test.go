package repository

import (
	"clean-code/model"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyStudents = []model.Student{
	{
		Id:       1,
		Name:     "Dummy name 1",
		Gender:   "M",
		Age:      1,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 1",
		Senior:   false,
	},
	{
		Id:       2,
		Name:     "Dummy name 2",
		Gender:   "F",
		Age:      2,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 2",
		Senior:   true,
	},
}

type StudentRepositoryTestSuite struct {
	suite.Suite
	mockResource *sqlx.DB
	mock         sqlmock.Sqlmock
}

func (suite *StudentRepositoryTestSuite) SetupTest() {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")
	suite.mockResource = sqlxDB
	suite.mock = mock
}

func (suite *StudentRepositoryTestSuite) TearDownTest() {
	suite.mockResource.Close()
}

func (suite *StudentRepositoryTestSuite) TestStudentRepository_GetAll() {
	rows := sqlmock.NewRows([]string{"id", "name", "gender", "age", "join_date", "id_card", "senior"})

	for _, d := range dummyStudents {
		rows.AddRow(d.Id, d.Name, d.Gender, d.Age, d.JoinDate, d.IdCard, d.Senior)
	}

	suite.mock.ExpectQuery("SELECT (.+) FROM m_student").WillReturnRows(rows)

	repo := NewStudentRepository(suite.mockResource)
	all, err := repo.GetAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))
	assert.Equal(suite.T(), 1, all[0].Id)
}

func (suite *StudentRepositoryTestSuite) TestStudentRepository_GetOneById() {
	rows := sqlmock.NewRows([]string{"id", "name", "gender", "age", "join_date", "id_card", "senior"})

	for _, d := range dummyStudents {
		rows.AddRow(d.Id, d.Name, d.Gender, d.Age, d.JoinDate, d.IdCard, d.Senior)
	}

	// SELECT (.+) FROM m_student WHERE id_card=\$1
	suite.mock.ExpectQuery("SELECT (.+) FROM m_student WHERE id_card=?").WithArgs("1").WillReturnRows(rows)

	repo := NewStudentRepository(suite.mockResource)
	one, err := repo.GetOneById("1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, one.Id)
}

func (suite *StudentRepositoryTestSuite) TestStudentRepository_GetOneByName() {
	rows := sqlmock.NewRows([]string{"id", "name", "gender", "age", "join_date", "id_card", "senior"})

	for _, d := range dummyStudents {
		rows.AddRow(d.Id, d.Name, d.Gender, d.Age, d.JoinDate, d.IdCard, d.Senior)
	}

	suite.mock.ExpectQuery("SELECT (.+) FROM M_STUDENT WHERE name like ?").WithArgs("Dummy name 1").WillReturnRows(rows)

	repo := NewStudentRepository(suite.mockResource)
	one, err := repo.GetOneByName("Dummy name 1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Dummy name 1", one[0].Name)
}

func TestStudentRepositoryTesSuite(t *testing.T) {
	suite.Run(t, new(StudentRepositoryTestSuite))
}
