package main

type Trec struct {
	v1 int
	vx struct {
		v2, v3 int
	}
	v4 int
}

func BanyakNilai(rec *Trec) {
	n := []int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4}

	for i, min := 0, n[0]; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
		rec.v1 = min
	}

	rec.vx.v2 = 0
	for _, v := range n {
		rec.vx.v2 += v
	}

	rec.vx.v3 = rec.vx.v2 / len(n)

	for i, max := 0, n[0]; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
		rec.v4 = max
	}
}
