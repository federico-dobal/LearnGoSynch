Problem statement

Use 2 functions printEven and printOdd to print numbers from 1 up to 10.
printEven: only print even numbers
printOdd: only print odd numbers

The	program	must	print	the	numbers	on	the	screen	in	increasing	order.	

Output expected:

0 1 2 3 4 5 6 7 8 9 

Solution without using printEven and printOdd numbers would be as follows:

func main() {
	fmt.Println("Init")
	for i:=0 ; i<10; i++ {
		fmt.Printf("%d ", i)
	} 
	fmt.Printf("\n")
}