package main

import "fmt"

func myAtoi(str string) int {
  retval := 0
  multiplier := 1
  index := 0
  numStart := false
  signSeen := false
  
  for index = 0; index < len(str); index++ {
    switch {
      case str[index] == ' ':
        if (signSeen) {
            //After a +/- sign no spaces are allowed
            return 0
        }
        //skip leading spaces
        continue
      case str[index] == '-':
        if (signSeen) {
          //we got a second - it is error - return 0
          return 0
        }
        multiplier = -1
        signSeen = true
      case str[index] == '+':
        if (signSeen) {
          return 0
        }
        signSeen = true
        //Multiplier is already 1

      case '0' <= str[index] && str[index] <= '9':
        numStart = true

      default:
        //Don't expect any other character
        return 0
    }
    if (numStart) {
      break
    }
  }

  //No numbers ?? So errors
  if (!numStart) {
    return 0
  }

STRING_ITER:
  for index < len(str) {
    switch {
      //Digit
      case '0' <= str[index] && str[index] <= '9':
        retval = retval*10+int(str[index]-'0')
        //If retval goes higher than MAX_INT (2147483647)
        // need to return 2147483647
        //Likewise for values below MIN_INT(-2147483648),
        // need to return -2147483648
        if ((retval < 0) || (retval > 2147483647)) {
            if (multiplier == -1) {
                retval = 2147483648
            } else {
                retval = 2147483647
            }
            break STRING_ITER
        }        
      //If we space, we terminate, if it is a 
      //negative number, it is invalid to get space
      //not exactly sure why, but that is the requirement
      case str[index] == ' ':
        if (multiplier == -1) {
            return 0
        }
        break STRING_ITER

      default:
        //no other character expected.. error
        break STRING_ITER
    }
    index ++
  }
  retval *= multiplier
  return (retval)
}

func main () {
  fmt.Println(myAtoi("18446744073709551617"))
  fmt.Println(myAtoi("9223372036854775809"))
  fmt.Println(myAtoi("  -0012a42"))
  fmt.Println(myAtoi("      -11919730356x"))
  fmt.Println(myAtoi("   +0 123"))
  //best way to check is to use it in arithmetic operation
  fmt.Println("101 + 1 = ", myAtoi("101")+myAtoi("1"))
  fmt.Println("+12345 + -20 = ", myAtoi(" +12345")+myAtoi("  -20"))

  // Negative cases
  fmt.Println("--101 = ", myAtoi("--101"))
  fmt.Println("10.1 = ", myAtoi("10.1"))
  fmt.Println(".   -101 = ", myAtoi(".   -101"))
  fmt.Println("-101AB = ", myAtoi("-101AB"))
  fmt.Println("10-1 = ", myAtoi("10-1"))
  fmt.Println("ABCD = ", myAtoi("ABCD"))
  fmt.Println("     = ", myAtoi("   "))  
  fmt.Println("= ", myAtoi("")) 
}
