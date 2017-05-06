/*
https://leetcode.com/problems/permutation-in-string

567. Permutation in String
--------------------------

Given two strings s1 and s2, write a function to return true if s2 contains the
permutation of s1. In other words, one of the first string's permutations is the
substring of the second string.
*/

package main

import "fmt"

func checkInclusion(s1b string, s2b string) bool {
  // to store the character frequency
  var freq [26]int

  //Stores the number of uniq characters seen
  uniqCh := 0

  //we are searching for s1 in s2 -
  //so s1 cannot be larger than s2
  if (len(s1b) > len(s2b)) {
    return false
  }

  //Get the frequencies in S1  
  for i:= 0; i < len(s1b); i++ {
    if (freq[s1b[i]-'a'] == 0) {
      uniqCh ++
    }
    freq[s1b[i]-'a']++
  }
  
  //Searching in the subject string (s2) - we'll do
  //it in two stages
  // #1: first substring - equal to the size of s1
  // #2: if no match, we'll do a sliding window to see
  // if any of the window of size len(s1) matches the
  // same frequency as s1

  // Do Step 1
  for i:= 0; i < len(s1b); i++ {
    freq[s2b[i]-'a']--
    //If we have matched all needed characters in 
    //the frequency array, then reduce the uniqCh to match 
    //This helps avoid having us walk the 26 character array
    if (freq[s2b[i]-'a'] == 0) {
      uniqCh--
    }
  }
  //If number of uniqCh left to match is zero, we have 
  //a match !!
  if (uniqCh == 0) {
    return true
  }

  //Start the window - from left --> right 
  left := 0
  right := len(s1b)

  //Loop until right is the end of the search string
  for right < len(s2b) {
    // move right to one right and 
    // update frequency
    freq[s2b[right]-'a']--
    if (freq[s2b[right]-'a'] == 0) {
      uniqCh--
    }
    // move left and here we'll increase frequency
    // since we are kicking it out of the window
    freq[s2b[left]-'a']++
    //uniqCh is to be restored back if freq[] was zero
    if (freq[s2b[left]-'a'] == 1) {
      uniqCh ++
    }
    //move the window to the right
    left++
    right++
    // if we have uniqCh is zero, we have 
    //matched.. NOTE you should not check 
    //until the entire window is shifted.
    if (uniqCh == 0) {
      return true
    }
  }
  //We'll get here only we are at the end of the string2 and
  //we haven't yet matched all uniq characters and their frequencies
  return false
}

func main () {
  //Few testcases
  fmt.Println("hello:ooolleoooleh", checkInclusion("hello", "ooolleoooleh"))
  fmt.Println("hello:ooolleooolleh", checkInclusion("hello", "ooolleooolleh"))
  fmt.Println("hello:ooolleooollehfea", checkInclusion("hello", "ooolleooollehfea"))
  fmt.Println("hello:ooolleooleolhfea", checkInclusion("hello", "ooolleooleolhfea"))
}
