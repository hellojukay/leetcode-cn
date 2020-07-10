#https://leetcode-cn.com/problems/counting-bits/
class Solution:
    def countBits(self, num: int) -> List[int]:
        result = []
        for i in range(num+1):
            count = 0
            while i != 0:
                if i % 2 == 1:
                    count = count + 1
                i = i >> 1
            result.append(count)
        return result