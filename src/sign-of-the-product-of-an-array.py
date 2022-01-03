class Solution:
    def arraySign(self, nums: List[int]) -> int:
        flag = 1
        for n in nums:
            if n == 0:
                flag = 0
                break
            if n > 0:
                continue
            if n < 0:
                flag = -1 * flag
        if flag == 0:
            return 0
        if flag > 0:
            return 1
        return -1
