package xs

import (
	"gopkg.in/go-playground/assert.v1"
	"strconv"
	"testing"
)

func TestGetTags(t *testing.T) {
	t.Parallel()
	type L struct {
		A int     `xs:"a"`
		B float64 `xs:"b:0.00%"`
		C string  `xs:"c:hyperlink"`
	}

	for i := 0; i <= 100; i++ {
		l := make([]L, 10, 10)

		for i := 0; i < 10; i++ {
			l[i] = L{
				A: i,
				B: float64(i) + 0.1,
				C: strconv.Itoa(i),
			}
		}

		//case1 l为结构体的值
		tags := GetTagInfo(l)
		assert.Equal(t, tags.M["a"].Omitempty, true)
		assert.Equal(t, tags.M["b"].Index, 1)
		assert.Equal(t, tags.M["c"].IsHyperlink, true)
		assert.Equal(t, tags.Headers, []string{"a", "b", "c"})

		//case2 l为结构体的指针
		tags2 := GetTagInfo(&l)
		assert.Equal(t, tags2.M["a"].Omitempty, true)
		assert.Equal(t, tags.M["b"].Index, 1)
		assert.Equal(t, tags.M["c"].IsHyperlink, true)
		assert.Equal(t, tags.Headers, []string{"a", "b", "c"})
	}

}
