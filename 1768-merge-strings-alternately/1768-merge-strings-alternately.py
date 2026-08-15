class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        total_len = len(word1) + len(word2)
        res = [''] * total_len

        j = 0

        for i in range(max(len(word1), len(word2))):
            if i < len(word1):
                res[j] = word1[i]
                j += 1

            if i < len(word2):
                res[j] = word2[i]
                j += 1

        return ''.join(res)
