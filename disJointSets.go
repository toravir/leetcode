/*

352. Data Stream as Disjoint Intervals

Given a data stream input of non-negative integers a1, a2, ..., an, ..., 
summarize the numbers seen so far as a list of disjoint intervals.

For example, suppose the integers from the data stream are 1, 3, 7, 2, 6, ..., then the summary will be:

[1, 1]
[1, 1], [3, 3]
[1, 1], [3, 3], [7, 7]
[1, 3], [7, 7]
[1, 3], [6, 7]

*/
package main

import "fmt"
 type Interval struct {
	   Start int
	   End   int
 }

type SummaryRanges struct {
	List []Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
    var retVal SummaryRanges
    retVal.List = make([]Interval, 0)
    return retVal
}

func (this *SummaryRanges) Addnum(val int)  {
    debug := false
    if (debug) {
        fmt.Println(val, "beginning:", this.List)
    }
    if (len(this.List) == 0) {
        //list is empty
        this.List = append(this.List, Interval{val, val})
    } else {
        start := 0
        end := len(this.List)-1
        if (this.List[start].Start > val) {
            //Is it the left edge
            if (this.List[start].Start-1 == val) {
                this.List[start].Start = val
            } else {
                this.List = append([]Interval{Interval{val, val}}, this.List...)
            }
        } else if (this.List[end].End < val) {
            //Is it the right edge
            if (this.List[end].End+1 == val) {
                this.List[end].End = val
            } else {
                this.List = append(this.List, Interval{val, val})
            }
        } else {
            //Somewhere in the middle - do binary search
            for start < end-1 {
                mid := (start + end)/2
                if (this.List[mid].Start < val) {
                    //move right
                    start = mid
                } else {
                    //move left
                    end = mid
                }
            }
            //Terminated when start & end are right next to each other

            if (this.List[start].End+1 == val) {
                //This is just to the right of start
                if (this.List[end].Start-1 == val) {
                    // this val is the missing number between start range and end range, so
                    //merge
                    this.List[end].Start = this.List[end-1].Start
                    this.List = append(this.List[:end-1], this.List[end:]...)
                } else {
                    // Just expand start range
                    this.List[start].End = val
                }
            } else if (this.List[end].Start -1 == val) {
                //This number is just to the left of end range - so
                //expand end range. We do not need to check for merge - since
                // it was done already
                this.List[end].Start = val
            } else {
                //This number is either included in the 'start'/'end' range or
                //need to be added as a new range
                if ((this.List[start].Start <= val && this.List[start].End >= val) ||
                    (this.List[end].Start <= val && this.List[end].End >= val)) {
                    //Do nothing - since the val is included in
                    //either start range or end range
                } else {
                    //This is a new range between start & end
                    front := this.List[:end]
                    tail := this.List[end:]
                    //Expand the array one more element
                    tail = append(tail, Interval{0,0})
                    //Shift the tail to include space for one more entry
                    copy(tail[1:], tail[0:])
                    tail[0] = Interval{val, val}
                    this.List = append(front, tail...)
                    if (debug) {
                        fmt.Println("final:", this.List)
                    }
                }
            }
        }
    }
}


func (this *SummaryRanges) Getintervals() []Interval {
    return this.List
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */

func main () {
    //TC #1 - cases with duplicate values, insert in the middle, edges, merging ranges
    obj := Constructor()
    obj.Addnum(6)
    fmt.Println(obj.Getintervals())
    obj.Addnum(6)
    fmt.Println(obj.Getintervals())
    obj.Addnum(0)
    fmt.Println(obj.Getintervals())
    obj.Addnum(4)
    fmt.Println(obj.Getintervals())
    obj.Addnum(8)
    fmt.Println(obj.Getintervals())
    obj.Addnum(7)
    fmt.Println(obj.Getintervals())
    obj.Addnum(6)
    fmt.Println(obj.Getintervals())
    obj.Addnum(4)
    fmt.Println(obj.Getintervals())
    obj.Addnum(5)
    fmt.Println(obj.Getintervals())

    //TC #2 - set of inserts followed by the reverse set of duplicate inserts
    obj = Constructor()
    obj.Addnum(3)
    fmt.Println(obj.Getintervals())
    obj.Addnum(1)
    fmt.Println(obj.Getintervals())
    obj.Addnum(7)
    fmt.Println(obj.Getintervals())
    obj.Addnum(2)
    fmt.Println(obj.Getintervals())
    obj.Addnum(6)
    fmt.Println(obj.Getintervals())
    obj.Addnum(6)
    fmt.Println(obj.Getintervals())
    obj.Addnum(2)
    fmt.Println(obj.Getintervals())
    obj.Addnum(7)
    fmt.Println(obj.Getintervals())
    obj.Addnum(1)
    fmt.Println(obj.Getintervals())
    obj.Addnum(3)
    fmt.Println(obj.Getintervals())

    //TC #3 Stress - multiple duplicates etc - hard to validate though..
    obj = Constructor()
    tc := []int {
    34,191,386,10,372,131,435,286,286,455,50,217,104,191,17,424,2,370,252,311,4,4,143,421,470,154,472,342,213,385,152,184,419,492,336,180,6,14,303,134,141,422,328,223,199,150,440,11,401,183,90,167,178,275,217,379,389,464,387,147,95,281,86,209,8,323,236,122,119,236,302,66,203,462,261,269,127,253,80,164,363,220,324,446,496,184,415,158,364,499,84,351,209,213,213,92,38,162,210,73,113,253,100,444,94,1,352,209,196,255,217,226,163,449,82,286,369,409,385,37,90,236,87,439,368,194,425,76,318,24,101,35,147,234,254,220,102,202,188,63,259,190,93,244,195,125,204,270,240,88,432,254,337,219,212,151,69,268,499,15,381,258,295,256,116,359,234,273,296,284,21,389,359,79,110,340,215,497,18,306,395,154,475,122,159,201,36,323,349,392,44,162,295,267,327,234,435,462,389,24,69,88,447,1,198,482,469,410,266,173,367,303,345,283,251,190,433,30,310,314,292,90,213,240,328,247,45,12,291,375,80,166,388,232,64,391,207,73,389,489,159,367,331,251,341,238,275,70,500,76,203,57,175,111,186,51,352,183,476,490,22,397,455,424,136,298,451,446,431,441,363,467,462,431,173,57,58,326,36,400,486,343,455,397,418,76,114,87,258,152,254,15,202,214,462,397,115,174,461,117,247,167,278,142,144,10,43,133,442,144,300,337,483,316,201,216,297,29,222,91,125,131,402,340,316,467,126,419,236,478,332,309,381,169,440,431,393,352,8,8,417,114,309,73,89,243,171,303,178,237,152,363,17,12,250,108,328,256,65,206,408,197,329,76,94,300,97,311,423,268,7,380,137,204,144,80,324,368,21,187,369,419,488,6,55,17,29,476,111,41,135,310,456,442,446,384,396,106,238,70,167,218,343,321,228,162,358,345,255,6,451,293,94,192,64,18,386,203,231,340,205,222,193,204,362,10,110,323,325,437,36,34,35,472,39,92,72,120,336,162,342,356,467,127,315,252,56,62,123,155,92,237,105,248,333,404,191,1,488,182,473,140,410,50,9,91,286,86,179,141,19,419,365,456,228,254,28,43,271,110,66,294,12,329,85,58,186,422,403,317,276,168,116,20,438,413,371,149,236,383,441,478,180,205,380,87,101,108,457,427,410,49,444,476,11,499,99,126,155,237,369,141,409,195,5,383,245,384,20,198,289,459,252,467,286,231,158,378,386,10,166,156,113,275,169,78,375,76,167,447,389,175,214,183,350,490,387,0,19,189,473,290,191,419,258,374,275,120,345,102,59,307,5,232,132,198,205,214,46,265,132,457,46,425,309,174,269,271,280,442,282,33,368,41,391,248,83,157,337,419,312,386,95,296,76,458,489,2,343,40,89,351,131,23,414,62,64,443,486,363,152,310,294 }
    for i := range(tc) {
        obj.Addnum(tc[i])
    }

    fmt.Println(obj.Getintervals())
}
