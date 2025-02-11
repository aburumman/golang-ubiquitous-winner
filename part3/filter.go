package main

// filter return a the value that a returned as true from the func 

func filter(pred func(int) bool,  values []int) []int {
	var out []int 

	for _ v := range values {
		if pred(v) {
			out = append(out, v)
		}
	}
return out
}

func isOdd(v int) bool {
	return v % 2 == 1
}

func main() {

	alist := []int{1, 2, 3, 4}
	filter(issOdd(25), alist))
}