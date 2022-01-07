from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        def mysort(list: List[int]) -> List[int]:
            length = len(list)
            if length <= 1:
                return list
            if length == 2:
                if list[0] < list[1]:
                    list[1], list[0] = list[0], list[1]
                    return list
                return list
            left = []
            mid = []
            right = []
            i = 0
            while i < length:
                if list[i] == list[int(length/2)]:
                    mid.append(list[i])
                    i = i + 1
                    continue
                if list[i] < list[int(length/2)]:
                    left.append(list[i])
                    i = i + 1
                    continue
                else:
                    right.append(list[i])
                i = i + 1
            return mysort(right) + mid + mysort(left)

        result = mysort(nums)
        return result[k - 1]



