# https://leetcode-cn.com/problems/reverse-bits/
class Solution:
    def reverseBits(self, n: int) -> int:
        result = 0
        for i in range(32):
            bit = n % 2
            n = n >> 1
            result = (result << 1) + bit
        return result

