impl Solution {
    pub fn check_perfect_number(num: i32) -> bool {
        let mut count = 0;
        for n in (1..(num/2)+1){
            if num % n == 0{
                count = count + n
            }
        }
        return count == num
    }
}
