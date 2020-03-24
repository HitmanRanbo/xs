package xs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

type User struct {
	Username string `xs:"Username"`
	Age      int    `xs:"Age"`
	Sex      string `xs:"Sex"`
}

func TestUnmarshalFromFile(t *testing.T) {
	var users = make([]User, 0)
	filePath := "example/test.xlsx"
	err := UnmarshalFromFile(filePath, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 4, "test.xlsx contain 4 col, and users does not contain")
	assert.Equal(t, users[3].Username, "Ann")
}

func TestUnmarshal(t *testing.T) {
	var users = make([]User, 0)
	filePath := "example/test.xlsx"
	body, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)
	err = Unmarshal(body, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 4, "test.xlsx contain 4 col, and users does not contain")
	assert.Equal(t, users[3].Username, "Ann")
}

type User2 struct {
	Username string  `xs:"Username"`
	Salary   float64 `xs:"Salary"`
}

func TestMarshal(t *testing.T) {
	users := []User{{"Karl", 25, "Male"}, {"Ann", 18, "Female"}}
	users2 := []User2{{"Karl", 8000.00}, {"Ann", 9999.50}}
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

	l := make([]L, length, length)

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

	newL0 := make([]L, 0)
	err = UnmarshalFromFile("tmp.xlsx", &newL0)
	assert.Error(t, err)

	newL1 := make([]L, 0)
	err = UnmarshalFromFile("tmp.xlsx", &newL0, &newL1)
	assert.NoError(t, err)
	assert.Equal(t, len(newL0), length)
	assert.Equal(t, len(newL1), length)
	assert.Equal(t, newL0[1].C, "http://www.github.com/1")
	assert.Equal(t, newL1[1].C, "http://www.github.com/1")

}
