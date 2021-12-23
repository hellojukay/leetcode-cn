# https://leetcode-cn.com/problems/ugly-number/
class Solution:
    def isUgly(self, n: int) -> bool:
        if n <= 0:
            return False
        if n == 1:
            return True
        if n % 2 == 0:
            return self.isUgly(int(n/2))
        if n % 3 == 0:
            return self.isUgly(int(n/3))
        if n % 5 == 0:
            return self.isUgly(int(n/5))
        return False
