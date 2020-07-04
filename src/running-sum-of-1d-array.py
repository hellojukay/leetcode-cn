class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        tmp = 0
        result = []
        for n in  nums:
            tmp = tmp + n
            result.append(tmp)
        return result

