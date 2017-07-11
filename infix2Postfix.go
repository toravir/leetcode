/*
   Infix to Postfix
   Given an infix expression. Conver the infix expression to postfix expression.
*/
package main

import "fmt"

type stackType struct {
    buf []byte
    //methods:
        //initStack()
        //push()
        //isEmpty()
        //pop()
        //printStack()
}

func (st *stackType) initStack () {
    st.buf = make([]byte,0)
}

func (st *stackType) push (val byte) {
    st.buf = append(st.buf, val)
}

// Caller shd check if the stack isEmpty()
// if empty and pop is called, it will return byte(0)
func (st *stackType) pop () byte {
    if (st.isEmpty()) {
        return byte(0)
    }
    length := len(st.buf)
    retval := st.buf[length-1]
    st.buf = st.buf[:length-1]
    return retval
}

func (st *stackType) peek () byte {
    if (st.isEmpty()) {
        return byte(0)
    }
    return st.buf[len(st.buf) - 1]
}

func (st *stackType) isEmpty () bool {
    if (len(st.buf)==0) {
        return true
    }
    return false
}

func (st *stackType) printStack () {
    for i,v := range st.buf {
        fmt.Printf("\nbuf[%d]:%c",i,v)
    }
}

func precedence (oper byte) int8 {
    switch (oper) {
        case '(': return (1)
        case '+': fallthrough
        case '-': return (2)
        case '*': fallthrough
        case '/': return (3)
        case '^': return (4)
    }
    return 0
}

func in2Postfix (str string) string {
    var st stackType
    st.initStack()
    pfix := ""
    for i := range str {
        //fmt.Printf("\nProcessing %c, pfix: %s", str[i], pfix)
        switch (str[i]) {
        case '*': fallthrough
        case '+': fallthrough
        case '-': fallthrough
        case '^': fallthrough
        case '/':
            for !st.isEmpty() && (precedence(st.peek()) >= precedence(str[i])) {
                pfix = string(append([]byte(pfix), st.pop()))
            }
            st.push(str[i])
        case '(':
            st.push(str[i])
        case ')':
            for !st.isEmpty() {
                st1 := st.pop()
                if (st1 == '(') {
                    break
                }
                pfix = string(append([]byte(pfix), st1))
            }
        default:
            pfix = string(append([]byte(pfix), str[i]))
        }
    }
    for !st.isEmpty() {
        pfix = string(append([]byte(pfix), st.pop()))
    }
    fmt.Println("")
    return pfix
}

func main () {
    fmt.Println(in2Postfix("a+b*(c^d-e)^(f+g*h)-i"))
    fmt.Println(in2Postfix("(a+b)*(c+d)"))
    fmt.Println(in2Postfix("p*q+x*y"))
    fmt.Println(in2Postfix("(a+b+e)*(c+d+f)"))
}
