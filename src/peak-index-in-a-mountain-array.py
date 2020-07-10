# https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
class Solution:
    def peakIndexInMountainArray(self, A: List[int]) -> int:
        index = 0
        m = A[0]
        for i in range(len(A)):
            if A[i] >= m :
                m = A[i]
                index = i
            else:
                break
        return index