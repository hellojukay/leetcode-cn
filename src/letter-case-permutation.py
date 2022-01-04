class Solution:
    def letterCasePermutation(self, s: str) -> List[str]:
        if not s:
            return []

        def dfs(word: str, i: int, result: List[str]) -> None:
            if i < 0:
                result.append("".join(word))
                return
            ch = word[i]
            if ch.isdigit():
                dfs(word, i - 1, result)
                return
            if ch.islower():
                dfs(word, i - 1, result)
                ch = ch.upper()
                word[i] = ch
                dfs(word, i - 1, result)
                return
            if ch.isupper():
                dfs(word, i - 1, result)
                ch = ch.lower()
                word[i] = ch
                dfs(word, i - 1, result)

        result = []
        dfs(list(s), len(s) - 1, result)
        return result
