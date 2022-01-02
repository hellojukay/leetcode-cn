class Solution:
    def isSameAfterReversals(self, num: int) -> bool:
        if num == 0:
            return True
        s = str(num)[::-1]
        s = s.lstrip("0")

        s2 = s[::-1]
        s2 = s2.lstrip("0")
        return s2 == str(num)
