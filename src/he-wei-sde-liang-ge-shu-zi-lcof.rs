use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::new();
        let mut a :i32 = 0;
        let mut b :i32 = 0;
        for i in nums.iter() {
            map.insert(*i,true);
        }
        for i in nums.iter() {
            let flag = match map.get(&(target-i)) { 
                Some(x) => true,
                _ => false,
            };
            if flag {
                if *i != target - *i  {
                    a = target-*i;
                    b = *i;
                    break;
                }
            }
        }
        let mut v = Vec::new();
        v.push(a);
        v.push(b);
        v
    }
}

