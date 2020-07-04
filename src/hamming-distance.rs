impl Solution {
    pub fn hamming_distance(x: i32, y: i32) -> i32 {
        let mut n: i32 = 0;
        let mut i = 0;
        let mut a = x;
        let mut b = y;
        while (i < 32) {
            if((a%2)!=(b%2)) {
                n = n + 1 ;
            }
            a = a / 2;
            b = b / 2;
            i = i + 1;
        }
        return n
    }
}
