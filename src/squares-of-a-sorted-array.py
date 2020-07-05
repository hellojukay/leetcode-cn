class Solution:
    def sortedSquares(self, A: List[int]) -> List[int]:
        result = []
        for n in A :
            result.append(n*n)
        result.sort()
        return result
