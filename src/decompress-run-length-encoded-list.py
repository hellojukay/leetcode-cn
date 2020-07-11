#https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        result = []
        for i in range(0,len(nums),2):
            for count in range(nums[i]):
                result.append(nums[i+1])
        return result

