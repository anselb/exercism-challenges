def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Strands are different lengths")

    hamming_distance = 0

    for dna_index in range(len(strand_a)):
        if strand_a[dna_index] != strand_b[dna_index]:
            hamming_distance += 1

    return hamming_distance
