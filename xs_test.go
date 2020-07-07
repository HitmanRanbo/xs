package xs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

type Student struct {
	Empty0 int
	Grade  int `xs:"Grade"`
	Empty1 string
	Class  int `xs:"Class"`
	Empty2 bool
}

type User struct {
	Empty0 int
	Student
	Username string `xs:"Username"`
	Age      int    `xs:"Age"`
	Empty1   string
	Sex      string `xs:"Sex"`
	Empty2   bool
}

func TestUnmarshalFromFile(t *testing.T) {
	//test function
	var users = make([]User, 0)
	filePath := "example/test.xlsx"
	err := UnmarshalFromFile(filePath, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 4, "test.xlsx contain 4 row, and users does not contain")
	assert.Equal(t, users[3].Username, "Ann")

	//test error
	var inter interface{}
	err = UnmarshalFromFile(filePath, inter)
	assert.Error(t, err)
}

func TestUnmarshal(t *testing.T) {
	var users = make([]User, 0)
	filePath := "example/test.xlsx"
	body, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)
	err = Unmarshal(body, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 4, "test.xlsx contain 4 row, and users does not contain")
	assert.Equal(t, users[3].Username, "Ann")
}

type User2 struct {
	Empty0   int
	Username string `xs:"Username"`
	Empty1   string
	Salary   float64 `xs:"Salary"`
	Empty2   bool
}

func TestMarshal(t *testing.T) {
	users := []User{{Student: Student{Grade: 6, Class: 5}, Username: "Karl", Age: 25, Sex: "Male"}, {Student: Student{Grade: 7, Class: 5}, Username: "Ann", Age: 19, Sex: "Female"}}
	users2 := []User2{{Username: "Karl", Salary: 8000.00}, {Username: "Ann", Salary: 9999.50}}
	body, err := Marshal(users, users2)
	assert.NoError(t, err)
	err = ioutil.WriteFile("user.xlsx", body, 06666)
	assert.NoError(t, err)
}

//test Marshal and UnmarshalFromFile
func Test_Xs(t *testing.T) {
	t.Parallel()
	type L struct {
		A int     `xs:"a"`
		B float64 `xs:"b:0.00%"`
		C string  `xs:"c:hyperlink"`
	}

	length := 30

	l := make([]L, length)

	for i := 0; i < length; i++ {
		l[i] = L{
			A: i,
			B: float64(i) + 0.1,
			C: "http://www.github.com/" + strconv.Itoa(i),
		}
	}

	body, err := Marshal(l, l)
	assert.NoError(t, err)

	err = ioutil.WriteFile("tmp.xlsx", body, 06666)
	assert.NoError(t, err)

	//the length of struct slice can be less than the num of sheets
	newL0 := make([]L, 0)
	err = UnmarshalFromFile("tmp.xlsx", &newL0)
	assert.NoError(t, err)

	newL1 := make([]L, 0)
	err = UnmarshalFromFile("tmp.xlsx", &newL0, &newL1)
	assert.NoError(t, err)
	assert.Equal(t, len(newL0), length)
	assert.Equal(t, len(newL1), length)
	assert.Equal(t, newL0[1].C, "http://www.github.com/1")
	assert.Equal(t, newL1[1].C, "http://www.github.com/1")

	//the length of struct slice should not be less than the num of sheets
	newL2 := make([]L, 0)
	err = UnmarshalFromFile("tmp.xlsx", &newL0, &newL1, &newL2)
	assert.Error(t, err)
}
