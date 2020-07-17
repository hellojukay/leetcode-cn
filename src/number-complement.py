class Solution:
    def findComplement(self, num: int) -> int:
        stack = []
        while num:
            c = num % 2
            stack.append(c)
            num = num >> 1
        result = 0
        while stack:
            c = stack.pop(-1)
            if c == 1:
                c = 0
            else:
                c = 1
            result = result * 2 + c
        return result

