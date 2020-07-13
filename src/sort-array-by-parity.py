class Solution:
    def sortArrayByParity(self, A: List[int]) -> List[int]:
        result = list()
        for n in A:
            if n % 2 == 0:
                result.append(n)
        for n in A:
            if n % 2 != 0:
                result.append(n)
        return result
