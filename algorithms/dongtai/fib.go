package dongtai

// Fib 斐波那锲
func Fib(n int) int {
	if n < 2{
		return n
	}
	a,b,c := 0,1,0
	for i:= 1;i<n;i++{
		c = a + b
		a,b = b,c
	}
	return c
}