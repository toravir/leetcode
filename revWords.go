/*
https://leetcode.com/problems/reverse-words-in-a-string-iii

557. Reverse Words in a String III
-----------------------------------

Given a string, you need to reverse the order of characters in each word within a
sentence while still preserving whitespace and initial word order.
*/

package main

import "fmt"

func reverseWords(st1 string) string {
  //lets do it in place - lets convert it to a
  //mutable byte array - add an extra 'space'
  //which is used by the logic to detect end of 
  //each word
  s := []byte(st1+" ")

  //Start index of the word
  start:=-1
  //end index of the current word
  end:=-1
  //keep a boolean to know if we are in the middle of a word
  word := false

  //Walk from start to end of the string
  for i:=0;i<len(s);i++ {
    //space means we may be at the end of the word
    if (s[i]==' ') {
      if (word) {
        end = i-1
        //fmt.Println("EOS: Start:", start, "end:", end)
        //Swap bytes starting at the ends to the middle
        for j:=0;j<(end-start)/2+1;j++ {
          t := s[start+j]
          s[start+j]=s[end-j]
          s[end-j]=t
          //fmt.Println("swapping:",start+j, start+end-j)
        }
        //we are out of a word
        word = false
      }
    } else {
      //If we see a non blank character, we are 
      //either starting a word or we are in 
      // the middle of one.
      if (word == false) {
        word = true
        start = i
      }
    }
  }
  //return the string with the last ' ' removed
  return (string(s[:len(s)-1]))
}

func main() {
    fmt.Println("abc defg  ij  kl:", reverseWords("abc defg  ij  kl"))
    fmt.Println("abc defg  ij  l:", reverseWords("abc defg  ij  l"))
    fmt.Println("abc defg  ij  :", reverseWords("abc defg  ij  "))
    fmt.Println("l:", reverseWords("l"))
    fmt.Println(":", reverseWords(""))    
}

