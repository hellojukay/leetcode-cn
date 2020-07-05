class Solution:
    def maximum(self, a: int, b: int) -> int:
        diff = abs((a - b) / 2)
        return int((a+b)/2 + diff)

