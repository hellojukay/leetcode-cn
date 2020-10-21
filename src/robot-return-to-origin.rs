impl Solution  {
    pub fn judge_circle(moves: String) -> bool {
        let mut L = 0;
        let mut R = 0;
        let mut U = 0;
        let mut D = 0;
        for ch in moves.chars() {
            match ch {
                'D' => D = D + 1,
                'U' => U = U + 1,
                'L' => L = L + 1,
                'R' => R = R + 1,
                _ => println!("error char"),
            }
        }
        return (L == R) && (U ==D);
    }
}

