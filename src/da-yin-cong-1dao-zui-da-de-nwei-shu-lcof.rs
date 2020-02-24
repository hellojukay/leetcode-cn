// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
impl Solution {
    pub fn print_numbers(n: i32) -> Vec<i32> {
        let max = match n {
            1 => 9,
            2 => 99,
            3 => 999,
            4 => 9999,
            5 => 99999,
            _ => 999999,
        };
        let mut vec = Vec::new();
        for n in 1..=max {
            vec.push(n);
        }
        vec
    }
}
