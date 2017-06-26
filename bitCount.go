/*
https://leetcode.com/problems/counting-bits

338. Counting Bits

Given a non negative integer number num. For every numbers i in the range 
0 ≤ i ≤ num calculate the number of 1's in their binary representation and 
return them as an array.

*/

package main

import "fmt"

func countBits(num int) []int {
    ret := make([]int, num+1)
    even := false
    ret[0]=0
    prevNumBits := 0
    for i:= 1; i <= num; i++ {
        //odd number means - we set one extra bit (lsb)
        if (!even) {
            prevNumBits ++
            ret[i] = prevNumBits
            even = !even
            continue
        }
        //We are becoming even...
        if (i & (i-1)==0) {
            //new number is a power of 2 - so we go back to 
            //only one bit set
            ret[i] = 1
            prevNumBits = 1
        } else {
          //fmt.Println("Computing num bits in:", i, "using: ", i>>1)
          ret[i]=ret[i>>1]
          prevNumBits = ret[i]
        }
        even = !even
    }
    return ret
}

func main () {
  fmt.Println(countBits(1))
  fmt.Println(countBits(20))
}
