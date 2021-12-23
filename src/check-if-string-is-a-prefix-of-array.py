class Solution:
    def isPrefixString(self, s: str, words: List[str]) -> bool:
        tmp = ''
        for ss in words:
            tmp = tmp + ss
            if s == tmp:
                return True
        return False
