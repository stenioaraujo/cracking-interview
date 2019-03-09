def main():
    print(one_away("pale", "ple"))
    print(one_away("pales", "pale"))
    print(one_away("pale", "bale"))
    print(one_away("pale", "bake"))
    
    print(one_away("abc", "ab"))
    print(one_away("ab", "abc"))
    print(one_away("", ""))
    print(one_away("pale", "pae"))
    print(one_away("pale", "pan"))

    print(one_away("abcd", "ab"))

def one_away(s1, s2):
    len_s1 = len(s1)
    len_s2 = len(s2)

    diffs = 0
    if len_s1 == len_s2:
        for i in range(len_s1):
            if s1[i] != s2[i]:
                diffs += 1
    elif abs(len_s1 - len_s2) == 1:
        longest, shortest = (s1, s2) if len_s1 > len_s2 else (s2, s1)
        i, j = 0, 0
        while i < len(longest) and j < len(shortest):
            if longest[i] != shortest[j]:
                diffs += 1
            else:
                j += 1
            i += 1
    else:
        return False
    
    return diffs <= 1

main()