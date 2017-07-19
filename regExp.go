/*
44. Wildcard Matching

Implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") ? false
isMatch("aa","aa") ? true
isMatch("aaa","aa") ? false
isMatch("aa", "*") ? true
isMatch("aa", "a*") ? true
isMatch("ab", "?*") ? true
isMatch("aab", "c*a*b") ? false

*/
package main

import "fmt"

/*
This is dynamic programming with recursion - using a cache 
to re-use previously computed values. The cache is populated
with the result of (i, j, result) - a key is derived from i,j
and the result is stored against that key.
*/
var cache map[int]bool

// Check if the cache has the value, if so, return that and 
// second return value tells if the cacheLookup was successful or not.
func checkCache(i int, j int) (bool, bool) {
  key := i*10000+j
  v,ok := cache[key]
  return v,ok
}

// Add to the cache
func addCache(i int, j int, b bool) {
  key := i*10000+j
  cache[key]=b
}


// Main function to do the matching
// This optimizes the pattern string:
// multiple consecutive '*' is converted to a single '*'
// also '*' is a really a special meaning here - it means '?*'
//* usually indicates the preceding character can occur zero or more times
// This function converts '*' to '?*'
func isMatch (str string, pat string) bool {
  cache = make(map[int]bool, 0)
  newpat := []rune{}
  for _,v := range(pat) {
    if (v == '*') {
      if (len(newpat) > 0 && newpat[len(newpat)-1]=='*') {
        continue
      }
      newpat = append(newpat, '?')
    }
    newpat = append(newpat, v)
  }
  np := string(newpat)
  //fmt.Println("newPat:", np)
  return checkMatch(str, np, len(str)-1, len(np)-1)
}


//Checks if the character matches the character in the pattern
//'ch' must match the pattern byte exactly - except
// when the pattern is '?' - which will match with any character
func isCharMatch (ch, pat byte) bool {
  if (pat == '?') {
    return true
  }
  if (ch == pat) {
    return true
  }
  return false
}

//Recursive function that will re-use any previously calculated
//values.
func checkMatch(str string, pat string, i int, j int) bool {
  //Case #1: if i == -1 and j == -1, means two null strings always match
  if (i < 0 && j < 0) {
    return true
  }
  //Case #2: if there is no pattern, and there is some string - then no match
  if (j < 0) {
    return false
  }
  v,ok := checkCache(i,j)
  if (ok) {
    //If we found this value in cache - reuse it
    return v
  }
  //Case #3:
  //When we are out of characters in the input but the pattern
  //has characters, we need to check what are the characters in the pattern
  if (i < 0) {
    if (pat[j]=='*') {
      q:= checkMatch(str, pat, i, j-2)
      addCache(i,j, q)
      return q
    }
    addCache(i,j, false)
    return false
  }

  //Case #4: - if the characters at i and j are matching,
  //Then rely on the match without these characters
  if (isCharMatch(str[i], pat[j])) {
    q := checkMatch(str, pat, i-1, j-1)
    addCache(i,j, q)
    return q
  }

  //Case #5: If pattern is '*', we need to check two sub-cases:
  // sub-case a: there are zero occurances - check if there is match with pattern[j-2] 
  //             -2 meaning 'skip the * and character before *'
  // sub-case b: If current character matches pattern[j-1], then remove that character and
  //             try the rest of pattern
  // The final result is OR operation of sub-case a and sub-case b.

  if (pat[j] == '*') {
    q1 := checkMatch(str, pat, i, j-2)
    if (isCharMatch(str[i], pat[j-1])) {
      q2 := checkMatch(str, pat, i-1, j)
      q1 = q1 || q2
    }
    addCache(i,j, q1)
    return (q1)
  }

  //Case #6: If none of the above matches, then it is FALSE.
  addCache(i,j, false)
  return false
}

func main () {
  fmt.Println(
    isMatch("abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab",
      "*aabb***aa**a******aa*"))
  fmt.Println(isMatch("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"))
}
