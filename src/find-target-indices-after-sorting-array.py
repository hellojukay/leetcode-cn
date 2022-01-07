class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        return [x for x, n in enumerate(sorted(nums)) if n == target]
