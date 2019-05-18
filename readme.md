# Coding problem: "time ranges subtraction"
Write a program that will subtract one list of time ranges from another. Formally: for two
lists of time ranges A and B, a time is in (A-B) if and only if it is part of A and not part of
B.

A time range has a start time and an end time. You can define times and time ranges
however you want (Unix timestamps, date/time objects in your preferred language, the
actual string “start-end”, etc).

Your solution shouldn’t rely on the granularity of the timestamps (so don’t, for example,
iterate over all the times in all the ranges and check to see if that time is “in”).

## Examples:
1. (9:00-10:00) “minus” (9:00-9:30) = (9:30-10:00)
2. (9:00-10:00) “minus” (9:00-10:00) = ()
3. (9:00-9:30) “minus” (9:30-15:00) = (9:00-9:30)
4. (9:00-9:30, 10:00-10:30) “minus” (9:15-10:15) = (9:00-9:15, 10:15-10:30)
5. (9:00-11:00, 13:00-15:00) “minus” (9:00-9:15, 10:00-10:15, 12:30-16:00)
= (9:15-10:00, 10:15-11:00)

# John Finnson's Go Implementation

- (Some of my golang terminology might be inaccurate. Some terms might stem from Python. Like "slice" vs "list".)
- (Moderate level of proficiency with golang)

## Setup
- Developed and run with: go 1.12.5

## To Run
- RUN `go run challenge.go`

## To Test
- RUN `go test`

## Assumptions
- Time ranges can be any time since UNIX Epoch (https://en.wikipedia.org/wiki/Unix_time)
- Dont need to protect functions with validations and assertion
- Phrase "granularity of the timestamps" refers to an approach where (2am-5am) would be broken down into 
(2am, 3am, 4am, 5am)

## Design Rational
- Core design philosophy is around combining all ranges (positive and negative) in a single list/slice. This slice can 
be used to find sets of time ranges that are unblocked by the negative ranges.
- Time is stored as golang time object. Unix Epoch time representation is used for comparisons.

### Semi-Visual example:
- rangeA - rangeB
- rangeA = [1+++4, 7++++++13] (1am-4am, 7am-1pm)
- rangeB = [2-3, 10-----15] (2am-3am, 10am-3pm)
- combined = [1+, 2-, -3, +4, 7+, 10-, +13, -15] (1am start range, 2am start block, 3am block end, 4am range end, 
7am range start, 10am block start, 1pm range end, 3pm block end)
- iterating through, sets found = [1+, 2-], [-3, +4], [7+, 10-] (1am-2am, 3am-4am, 7am-10am)

1. Time ranges are broken down into individual times, tagged with a "type", then combined into a single slice.
2. The slice of time ranges are then sorted by time (earliest first)
3. The sorted slice is iterated over, creating ranges as spans of "unblocked" time are found

### Types Terminology
- "Rs" == "Range start" == start of "positive" range from range A
- "Re" == "Range end" == end of "positive" range from range A
- "Bs" == "Block start" == start of "negative" range from range B
- "Be" == "Block end" == end of "negative" range from range B

## Complexity
- Total time complexity:
    - Worst case: O(n * log^2(n))
- golang sort.Slice is O(n * log^2(n)) (https://golang.org/src/sort/sort.go)
- extractTimeRanges is O(n)

## TODO (if this was production code)
- Add assertions and error handling
- Consider using a testing framework like Ginkgo (https://github.com/onsi/ginkgo)
- Clean up the struct, field, and fn names a bit. (naming is hard!)
