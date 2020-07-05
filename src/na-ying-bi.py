class Solution:
    def minCount(self, coins: List[int]) -> int:
        result = 0
        for n in coins :
            result = result + int(n/2)
            result = result + (n % 2)
        return result
