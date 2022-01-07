from typing import List


class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        def mySort(nums: List[int], m: dict[int, int]) -> List[int]:
            length = len(nums)
            if length <= 1:
                return nums
            if length == 2:
                if m.get(nums[0], -1) > m.get(nums[1], -1):
                    nums[0], nums[1] = nums[1], nums[0]
                return nums
            left = []
            right = []
            l = []
            mid = int(length / 2)
            i = 0
            while i < length:
                if m.get(nums[i], -1) == m.get(nums[mid], -1):
                    l.append(nums[i])
                    i = i + 1
                    continue
                if m.get(nums[i], -1) < m.get(nums[mid], -1):
                    left.append(nums[i])
                else:
                    right.append(nums[i])
                i = i + 1
            return mySort(left, m) + l + mySort(right, m)

        m = {}
        i = 0
        while i < len(arr2):
            m[arr2[i]] = i
            i = i + 1
        prefix = []
        suffix = []
        i = 0
        while i < len(arr1):
            if m.get(arr1[i], -1) >= 0:
                prefix.append(arr1[i])
            else:
                suffix.append(arr1[i])
            i = i + 1
        return mySort(prefix, m) + sorted(suffix)
