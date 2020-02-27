impl Solution {
    // https://leetcode-cn.com/problems/container-with-most-water/
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut  max = 0;
        for i in 0..height.len() {
            for j in (i+1)..height.len() {
                let mut area = 0;
                if height[j] > height[i] {
                    area = (j-i) as i32 *height[i];
                }else {
                    area = (j-i) as i32 *height[j];
                }
                if area > max {
                    max = area;
                }
            }
        }
        max
    }
}

