package URLShortener

import "math"

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62

//Function encodes the given database ID to a base62 string
func ToBase62(num int) string {
	r := num % b
	res := string(base[r])
	div := num / b
	q := int(math.Floor(float64(div)))

	for q != 0{
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[(int(r))]) + res
	}
	return string(res)
}
