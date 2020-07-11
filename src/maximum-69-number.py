# https://leetcode-cn.com/problems/maximum-69-number/
class Solution:
    def maximum69Number (self, num: int) -> int:
        result = 0
        mixed = False
        for ch in str(num):
            if not mixed and ch =='6':
                mixed = True
                ch = '9'
            result = result * 10 +  int(ch)
        return result


