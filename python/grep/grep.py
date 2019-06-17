def grep(pattern, flags, files):
    # Set the pattern_in_str parameters based on the flags
    check_case = True
    match = True
    full_line = False
    if pattern_in_str("-i", flags):
        check_case = False
    if pattern_in_str("-v", flags):
        match = False
    if pattern_in_str("-x", flags):
        full_line = True

    # Initialize return value
    found_lines = ""
    # Iterate through all the lines in each file, for all the passed in files
    for file in files:
        with open(file) as f:
            for line_index, line in enumerate(f.readlines()):
                # If there is an "-l" flag, only return the filename
                if pattern_in_str(pattern, line, check_case, match, full_line) and pattern_in_str("-l", flags):
                    found_lines += "{}\n".format(file)
                    # Add the text file name once
                    break
                # If there is no "-l" flag, return the line text
                elif pattern_in_str(pattern, line, check_case, match, full_line):
                    # If there is more than one text file, specify the filename
                    if len(files) > 1:
                        found_lines += "{}:".format(file)
                    # If there is an "-n" flag, specify the line index
                    if pattern_in_str("-n", flags):
                        # Start at index 1
                        found_lines += "{}:".format(line_index + 1)
                    found_lines += line

    return found_lines


def pattern_in_str(pattern, text, check_case=True, match=True, full_line=False):
    # If case does not matter, set the casing for pattern and text to lowercase
    if not check_case:
        pattern = pattern.lower()
        text = text.lower()
    # Iterate through each possible index in the text
    for str_ind in range(len(text) - len(pattern) + 1):
        # If matching the full line, and the text character to match the
        # first pattern character is also not the first one, it is not a match
        if full_line and str_ind != 0:
            return not match
        # Check if each character matches the first letter in the pattern
        if text[str_ind] == pattern[0]:
            # If there is, check every character in the pattern
            for pattern_ind in range(len(pattern)):
                # Iterate throuh the pattern and text to check for match
                if pattern[pattern_ind] != text[pattern_ind + str_ind]:
                    # If one character is off, stop checking
                    break
                # If whole pattern was cheked with no break, it matched
                if pattern_ind == len(pattern) - 1:
                    return match
    # If all characters have been checked, and nothing returns, the pattern is
    # not in the text
    return not match
