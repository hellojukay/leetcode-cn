// https://leetcode-cn.com/problems/check-permutation-lcci/
impl Solution {
    pub fn check_permutation(s1: String, s2: String) -> bool {
        let mut result = true;
        if s1.len() != s2.len() {
            result =false;
            return result
        }
        let mut v1 = Vec::new();
        let mut v2 = Vec::new();
        for ch in s1.chars() {
            v1.push(ch);
        }
        for ch in s2.chars() {
            v2.push(ch);
        }
        v1.sort();
        v2.sort();
        for i in 0..v1.len() {
            if v1[i] != v2[i] {
                result = false;
                break
            }
        }
        result
    }
}
