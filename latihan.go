package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, jumUang int
	fmt.Scan(&n)
	jumUang = 0
	for i = 1; i <= n; i++{
		fmt.Scan(&A[i])
		if i % 2 != 0{
			jumUang += A[i]
		}
	}
	fmt.Print(jumUang)
}	