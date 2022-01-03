class Solution:
    def twoOutOfThree(
        self, nums1: List[int], nums2: List[int], nums3: List[int]
    ) -> List[int]:
        s1 = set(nums1)
        s2 = set(nums2)
        s3 = set(nums3)
        result = set()
        for n in s1:
            if n in s2 or n in s3:
                result.add(n)
        for n in s2:
            if n in s1 or n in s3:
                result.add(n)
        for n in s3:
            if n in s1 or n in s2:
                result.add(n)
        return list(result)
