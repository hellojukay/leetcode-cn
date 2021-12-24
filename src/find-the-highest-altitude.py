class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        m = 0
        result = 0
        for n in gain:
            result = result + n
            if result > m:
                m = result
        return m
