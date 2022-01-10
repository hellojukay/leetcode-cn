class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        nums = range(0,n+1)
        return sorted(nums,key=str)
