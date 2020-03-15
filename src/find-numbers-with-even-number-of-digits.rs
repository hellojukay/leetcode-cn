impl Solution {
    pub fn find_numbers(nums: Vec<i32>) -> i32 {
        let mut sum = 0;
        for i in nums.iter() {
            if Solution::check(*i)%2 ==0 {
                sum = sum + 1;
            }
        }
        sum
    }
    pub fn check(n :i32) -> i32 {
        if n < 10 {
            return 1
        }
        if n < 100 {
            return 2
        }
        if n < 1000 {
            return 3
        }
        if n < 1000 {
            return 4
        }
        if n < 10000 {
            return 5
        }
        return 6
    }
}
