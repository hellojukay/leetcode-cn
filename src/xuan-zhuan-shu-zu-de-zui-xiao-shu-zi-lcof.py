class Solution:
    def minArray(self, numbers: List[int]) -> int:
        m = numbers[0]
        for n in numbers:
            if n < m:
                m = n
                break
        return m
