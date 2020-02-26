impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut stack = Vec::new();
        for i in nums.iter() {
            let length = stack.len();
            if length  == 0 {
                stack.push(*i);
                continue;
            }
            if *i != stack[length-1] {
                stack.pop();
            }else {
                stack.push(*i);
            }
        }
        stack[0]
    }
}
