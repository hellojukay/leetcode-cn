class Solution:
    def createTargetArray(self, nums: List[int], index: List[int]) -> List[int]:
        result = []
        for i in range(len(index)):
            result.insert(index[i],nums[i])
        return result
