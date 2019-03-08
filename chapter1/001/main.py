def unique_map(string):
    m = {}
    for c in string:
        if m.get(c):
            return False
        m[c] = True
    return True


def unique_inplace(string):
    for i in range(len(string)):
        for j in range(i+1, len(string)):
            if string[i] == string[j]:
                return False
    return True


print(unique_map("abc"))
print(unique_map("stẽnio"))
print(unique_map("abba"))

print(unique_inplace("abc"))
print(unique_inplace("stẽnio"))
print(unique_inplace("abba"))