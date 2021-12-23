from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        length = len(nums1) + len(nums2)
        # 偶数个
        n = 0
        result = []
        while True:
            if n > (length/2):
                break
            if len(nums1) > 0:
                if len(nums2) > 0:
                    if nums1[0] < nums2[0]:
                        result.append(nums1[0])
                        n += 1
                        nums1 = nums1[1:]
                        continue
                    else:
                        result.append(nums2[0])
                        n += 1
                        nums2 = nums2[1:]
                        continue
                else:
                    result.append(nums1[0])
                    n += 1
                    nums1 = nums1[1:]
                    continue
            else:
                result.append(nums2[0])
                n += 1
                nums2 = nums2[1:]
                continue
        if length % 2 == 0:
            return (result[-1+int((length/2))] + result[int(length/2)]) / 2.0
        else:
            return result[int((length-1)/2)]


s = Solution()
print(s.findMedianSortedArrays([], [1]))
