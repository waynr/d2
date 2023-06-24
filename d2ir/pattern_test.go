package d2ir_test

import (
	"testing"

	"oss.terrastruct.com/util-go/assert"
)

func testCompilePatterns(t *testing.T) {
	t.Parallel()

	tca := []testCase{
		{
			name: "escaped",
			run: func(t testing.TB) {
				m, err := compile(t, `animal: meow
action: yes
a\*: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 3, 0, nil, "")
				assertQuery(t, m, 0, 0, "meow", "animal")
				assertQuery(t, m, 0, 0, "yes", "action")
				assertQuery(t, m, 0, 0, "globbed", `a\*`)
			},
		},
		{
			name: "prefix",
			run: func(t testing.TB) {
				m, err := compile(t, `animal: meow
action: yes
a*: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "animal")
				assertQuery(t, m, 0, 0, "globbed", "action")
			},
		},
		{
			name: "suffix",
			run: func(t testing.TB) {
				m, err := compile(t, `animal: meow
jingle: loud
*l: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "animal")
				assertQuery(t, m, 0, 0, "globbed", "jingle")
			},
		},
		{
			name: "prefix-suffix",
			run: func(t testing.TB) {
				m, err := compile(t, `tinker: meow
thinker: yes
t*r: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "tinker")
				assertQuery(t, m, 0, 0, "globbed", "thinker")
			},
		},
		{
			name: "prefix-suffix/2",
			run: func(t testing.TB) {
				m, err := compile(t, `tinker: meow
thinker: yes
t*ink*r: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "tinker")
				assertQuery(t, m, 0, 0, "globbed", "thinker")
			},
		},
		{
			name: "prefix-suffix/3",
			run: func(t testing.TB) {
				m, err := compile(t, `tinkertinker: meow
thinkerthinker: yes
t*ink*r*t*inke*: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "tinkertinker")
				assertQuery(t, m, 0, 0, "globbed", "thinkerthinker")
			},
		},
		{
			name: "nested/prefix-suffix/3",
			run: func(t testing.TB) {
				m, err := compile(t, `animate.constant.tinkertinker: meow
astronaut.constant.thinkerthinker: yes
a*n*t*.constant.t*ink*r*t*inke*: globbed`)
				assert.Success(t, err)
				assertQuery(t, m, 6, 0, nil, "")
				assertQuery(t, m, 0, 0, "globbed", "animate.constant.tinkertinker")
				assertQuery(t, m, 0, 0, "globbed", "astronaut.constant.thinkerthinker")
			},
		},
		{
			name: "edge/1",
			run: func(t testing.TB) {
				m, err := compile(t, `animate
animal
an* -> an*`)
				assert.Success(t, err)
				assertQuery(t, m, 2, 4, nil, "")
				assertQuery(t, m, 0, 0, nil, "(animate -> animal)[0]")
				assertQuery(t, m, 0, 0, nil, "(animal -> animal)[0]")
			},
		},
		{
			name: "edge/2",
			run: func(t testing.TB) {
				m, err := compile(t, `shared.animate
shared.animal
sh*.(an* -> an*)`)
				assert.Success(t, err)
				assertQuery(t, m, 3, 4, nil, "")
				assertQuery(t, m, 2, 4, nil, "shared")
				assertQuery(t, m, 0, 0, nil, "shared.(animate -> animal)[0]")
				assertQuery(t, m, 0, 0, nil, "shared.(animal -> animal)[0]")
			},
		},
		{
			name: "edge/3",
			run: func(t testing.TB) {
				m, err := compile(t, `shared.animate
shared.animal
sh*.an* -> sh*.an*`)
				assert.Success(t, err)
				assertQuery(t, m, 3, 4, nil, "")
				assertQuery(t, m, 2, 4, nil, "shared")
				assertQuery(t, m, 0, 0, nil, "shared.(animate -> animal)[0]")
				assertQuery(t, m, 0, 0, nil, "shared.(animal -> animal)[0]")
			},
		},
		{
			name: "double-glob",
			run: func(t testing.TB) {
				m, err := compile(t, `shared.animate
shared.animal
**.style.fill: red`)
				assert.Success(t, err)
				assertQuery(t, m, 9, 0, nil, "")
				assertQuery(t, m, 8, 0, nil, "shared")
				assertQuery(t, m, 2, 0, nil, "shared.style")
				assertQuery(t, m, 2, 0, nil, "shared.animate")
				assertQuery(t, m, 2, 0, nil, "shared.animal")
			},
		},
	}

	runa(t, tca)

	t.Run("errors", func(t *testing.T) {
		tca := []testCase{}
		runa(t, tca)
	})
}
