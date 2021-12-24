class Solution:
    def countOne(self, n: int) -> int:
        sum = 0
        while n > 0:
            if n % 2 == 1:
                sum = sum + 1
            n = int(n / 2)
        return sum

    def countBits(self, n: int) -> List[int]:
        return [self.countOne(x) for x in range(0, n+1)]
