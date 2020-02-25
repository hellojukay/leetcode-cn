// https://leetcode-cn.com/problems/fibonacci-number/
impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n <= 1 {
            return n
        }
        let mut first = 0;
        let mut second = 1;
        let mut sum = 0;

        for x in 0..n-1 {
            sum = first + second;
            first = second;
            second = sum;
        }
        second
    }
}

