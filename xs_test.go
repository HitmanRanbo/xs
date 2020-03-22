package xs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

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
			C: "http://www.baidu.com/" + strconv.Itoa(i),
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
	assert.Equal(t, len(newL0), length+1)
	assert.Equal(t, len(newL1), length+1)
	assert.Equal(t, newL0[1].C, "http://www.baidu.com/0")
	assert.Equal(t, newL1[1].C, "http://www.baidu.com/0")

}
