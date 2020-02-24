// https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let mut count = 0 as i32;
        for v in grid  {
            let s = v.len();
            for j in 0..s  {
                if v[j] < 0 {
                    count = count + ((s - j) as i32);
                    break;
                }
            }
        }
        count
    }
}

