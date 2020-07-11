#https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
class Solution:
    def sumZero(self, n: int) -> List[int]:
        result = []
        if n % 2 == 1:
            result.append(0)
            n = (n-1)/2
        else:
            n = n / 2
        for i in range(1,n+1):
            result.append(i)
            result.append(i*-1)
        return result

