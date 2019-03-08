def isPermutation(s1, s2):
	chars = {}

	for c in s1:
	    chars[c] = chars.get(c, 0) + 1

	for c in s2:
	    chars[c] = chars.get(c, 0) - 1

	for _, v in chars.items():
		if v != 0:
			return False

	return True

print(isPermutation("", ""))
print(isPermutation("abba", "aabb"))
print(isPermutation("stẽnio", "ẽinost"))
print(isPermutation("arara", "aarra"))

print(isPermutation("", "aarra"))
print(isPermutation("arara", "arera"))