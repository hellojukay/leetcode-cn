// https://leetcode-cn.com/problems/unique-number-of-occurrences/
use std::collections::HashMap;
use std::collections::HashSet;


impl Solution {
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        let mut m: HashMap<&i32,i32>  = HashMap::new();
        for it in &arr {
            let count = match m.get(&it) {
                Some(x) => x+1 as i32,
                _ => 1 as i32,
            };
            m.insert(it,count);
        }
        let size = m.len();
        let mut set  = HashSet::new();
        for _, v in &m {
            set.insert(*v);
        }
        return size == set.len()
    }
}
