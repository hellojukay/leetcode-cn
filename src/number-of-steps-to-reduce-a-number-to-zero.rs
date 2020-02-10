
impl Solution {
    pub fn number_of_steps (num: i32) -> i32 {
        let mut count = 0;
        let mut x = num;
        let y = loop {
            if x == 0 {
                break count;
            }
            if x % 2 == 0 {
                x = x /2;
                count += 1;
            }else {
                x = x -1;
                count += 1;
            }
        };
        y
    }
}

