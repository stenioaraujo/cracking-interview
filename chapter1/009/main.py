def main():
	print(rotate("waterbottle", "erbottlewat"))

	print(rotate("stenio", "stenio"))
	print(rotate("stenioo", "ostenio"))
	print(rotate("stenioo", "enioest"))

def rotate(original, rotation):
    if len(original) == len(rotation):
        return original in (rotation + rotation)
    
    return False

main()