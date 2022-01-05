class Solution:
    def modifyString(self, s: str) -> str:
        def get_char(without: List[str]) -> str:
            for c in string.ascii_letters:
                if c not in without:
                    return c
            return ""
        i = 0
        pre = ''
        result = ""
        while i < len(s):
            if s[i].isalpha():
                result = result + s[i]
                i = i + 1
                continue
            else:
                without = []
                if i > 0:
                    without.append(result[i-1])
                if i + 1 < len(s):
                    without.append(s[i+1])
                c = get_char(without)
                result = result + c
                i = i + 1
        return result
