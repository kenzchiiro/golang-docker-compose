
package solution


import (
	"fmt"
)

func PatternOne(N int) {
	
        for i := 0; i < N; i++ {
	   tmp:=0
	   for j := 0; j < (N*2)-1; j++ {
		if j < i {
		   fmt.Print("O")
		}else if i < (N*2)-1 - j {
		   if tmp == 0 {
			fmt.Print("X")
			tmp++
		   }else{
			fmt.Print("O")
			tmp--
		   }
		}else{
		   fmt.Print("O")
		}
	   }
            fmt.Println("");
        }
        for i := N-1; i > 0; i-- {
	   tmp:=0
	   for j := 0; j < (N*2)-1; j++ {
		if i-1>j{
		   fmt.Print("O")
		}else if i < (N*2) - j {
		   if tmp == 0 {
			fmt.Print("X")
			tmp++
		   }else{
			fmt.Print("O")
			tmp--
		   }
		}else{
		   fmt.Print("O")
		}
	   }
            fmt.Println("");
        }
	
}

func PatternTwo(N int) {	
	K:=0
	if N%2 == 1{
		K = N/2 +1
	}else{
		K = N/2
	}
	N=N/2
	// devide pattern to four part
	// upper half of the pattern
	for i := 0; i < K; i++ {
		// upper left
		for j := 0; j < K; j++ {
			if i >= j {  
				fmt.Print("X")
			}else{
				fmt.Print(" ")
			}
		}
		// upper right
		for j := 0; j < N; j++ {
			if i >= N-1 - j {  
				fmt.Print("X")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println("");
	}
	// bottom half of the pattern
	for i := 0; i < N; i++ {
		// bottom left
		for j := 0; j < K; j++ {
			if i + j <= N - 1 {  
				fmt.Print("X")
			}else{
				fmt.Print(" ")
			}
		}
		// bottom right
		for j := 0; j < N; j++ {
			if i <= j {  
				fmt.Print("X")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println("");
	}
}


func PatternThree(N int) {

    for i:= 0;i < N;i++{
        for j:= 0;j < N;j++{

            // reflect (i, j) to the top left quadrant as (a, b)
            a:= j
            b:= i
            if a >= N / 2 {
                a = N - a - 1
            }
            if b >= N / 2 {
                b = N - b - 1
            }

            // calculate distance from center ring
            u:= (a - N / 2)
            if (u < 0) {
                u = u * -1
            }
            v:= (b - N / 2)
            if (v < 0) {
                v = v * -1
            }
            d:= 0
            if u > v {
                d = u
            } else {
                d = v
            }
            L:= N / 2
            if (N % 3 == 0) && N>3{
                L--
            }

            // fix the top-left-to-bottom-right diagonal
            if (i == j + 1 && i <= L) {
                d++
            }
            if (d + N / 2) % 2 == 0 || N <= 2{
                fmt.Print("X");
            } else {
                fmt.Print("Y");
            }
        }
        fmt.Println("");
    }
}
