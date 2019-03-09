def main():
	print(is_palindrome_permutation("Tact Coa"))
	print(is_palindrome_permutation("arara"))
	print(is_palindrome_permutation("Abba"))
	print(is_palindrome_permutation("obo"))
	print(is_palindrome_permutation(""))
	print(is_palindrome_permutation("a"))

	print(is_palindrome_permutation("apex"))
	print(is_palindrome_permutation("stenio"))
	print(is_palindrome_permutation("concatenar"))
	print(is_palindrome_permutation("pararalelepipedo"))

def is_palindrome_permutation(string):
    string = string.lower()
    string = string.replace(" ", "")

    gates = {}
    for c in string:
        gates[c] = not gates.get(c, False)
    
    opens = 0
    for gate_open in gates.values():
        if gate_open:
            opens += 1
    
    if opens > 1:
        return False
    
    return True

main()