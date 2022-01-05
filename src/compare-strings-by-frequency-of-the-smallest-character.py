class Solution:
    def numSmallerByFrequency(self, queries: List[str], words: List[str]) -> List[int]:
        def count(word: str) -> int:
            for c in string.ascii_letters:
                if c in word:
                    return word.count(c)
            return 0
        nums = [count(word) for word in words]
        result = []
        for word in queries:
            n = count(word)
            c = len([c for c in nums if n < c])
            result.append(c)
        return result
