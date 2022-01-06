class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        def dfs(nums1: List[int], current: List[int],result: List[List[int]]) -> None:
            if len(nums1) == 0:
                result.append(current)
                return
            i = 0
            while i < len(nums1):
                tmp = nums1.copy()
                tmp2 = current.copy()
                tmp2.append(tmp.pop(i))
                dfs(tmp,tmp2,result)
                i = i + 1
        result = []
        dfs(nums,[],result)
        return result
