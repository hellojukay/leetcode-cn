class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        def check(word: str, chars: str) -> bool:
            m = {}
            for ch in chars:
                n = m.get(ch, 0)
                m[ch] = n + 1
            for ch in word:
                n = m.get(ch, 0)
                if n == 0:
                    return False
                else:
                    n = n - 1
                    m[ch] = n
            return True

        return sum([len(s) for s in words if check(s, chars)])
