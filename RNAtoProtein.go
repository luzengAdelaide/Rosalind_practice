package main

//no imports!!!

type codon [3]byte

func asCodon(s string) codon {
	var c codon
	if len(s) != len(c) {
		panic("bad codon")
	}
	copy(c[:], s) //c[:] changes this into a slice of c
	return c
}

var code = map[codon]byte{
	asCodon("UUC"): 'F', asCodon("CUC"): 'L', asCodon("AUC"): 'I', asCodon("GUC"): 'V',
	asCodon("UUU"): 'F', asCodon("CUU"): 'L', asCodon("AUU"): 'I', asCodon("GUU"): 'V',
	asCodon("UUA"): 'L', asCodon("CUA"): 'L', asCodon("AUA"): 'I', asCodon("GUA"): 'V',
	asCodon("UUG"): 'L', asCodon("CUG"): 'L', asCodon("AUG"): 'M', asCodon("GUG"): 'V',
	asCodon("UCU"): 'S', asCodon("CCU"): 'P', asCodon("ACU"): 'T', asCodon("GCU"): 'A',
	asCodon("UCC"): 'S', asCodon("CCC"): 'P', asCodon("ACC"): 'T', asCodon("GCC"): 'A',
	asCodon("UCA"): 'S', asCodon("CCA"): 'P', asCodon("ACA"): 'T', asCodon("GCA"): 'A',
	asCodon("UCG"): 'S', asCodon("CCG"): 'P', asCodon("ACG"): 'T', asCodon("GCG"): 'A',
	asCodon("UAU"): 'Y', asCodon("CAU"): 'H', asCodon("AAU"): 'N', asCodon("GAU"): 'D',
	asCodon("UAC"): 'Y', asCodon("CAC"): 'H', asCodon("AAC"): 'N', asCodon("GAC"): 'D',
	asCodon("UAA"): '*', asCodon("CAA"): 'Q', asCodon("AAA"): 'K', asCodon("GAA"): 'E',
	asCodon("UAG"): '*', asCodon("CAG"): 'Q', asCodon("AAG"): 'K', asCodon("GAG"): 'E',
	asCodon("UGU"): 'C', asCodon("CGU"): 'R', asCodon("AGU"): 'S', asCodon("GGU"): 'G',
	asCodon("UGC"): 'C', asCodon("CGC"): 'R', asCodon("AGC"): 'S', asCodon("GGC"): 'G',
	asCodon("UGA"): '*', asCodon("CGA"): 'R', asCodon("AGA"): 'R', asCodon("GGA"): 'G',
	asCodon("UGG"): 'W', asCodon("CGG"): 'R', asCodon("AGG"): 'R', asCodon("GGG"): 'G',
}

func translate(rna []byte) []byte {
	protein := make([]byte, len(rna)/3)
	// could also write this as: for i = 0; i < len(rna); i+=3 {}
	var c codon
	for i := range protein {
		copy(c[:], rna[i*3:(i+1)*3])
		protein[i] = code[c]
		if protein[i] == '*' {
			return protein[:i] //slice it before the stop codon
		}
	}
	return protein
}

func main() {
	println(string(translate([]byte("AUGCCAGGGGAUAUACGCAGCAGUAUCGCCGCGUCGCACACUCUCGGCUUUGUAUAUGCACUGUGGGCCGUUUUACUACACGUCGCAGUAGGGAAUACGGAGGAUCUGCACUGCAAUACGACAGCACGUAGGAGGUUUGUGGCAACAGUUGACCAUUUCGCCGCUUCCCCUCUAUCGUUCCAUAGAGCUGGAUUCCUUAUCCUUGGGCGUGGACACUCCUAUCACUCGGAAGGGGCUAAAGCCGCGCCUUGCGAACUAAGAACAGGUAUUUAUUUUGGACCGCGAUCGUGCGUCUGCACAGUGUCGGAUCGUAGUACACCAGCAUCUUCGCUUCUCUCAGGUUACGUUCCGCUGCGCACACCCAGAUAUAGAGAAUUUCCCGGCCGAUUACGGUGCUAUGGAUGCCCUGGUCGCAAGCUCAAUAAGACACUCGCCCUGUACCCAAUGCUGGUAGGCAUAUUGGACUCACACGCGGGAACGUGUCUGGUGGGCUCAGAUCCUGCUGUACUAGCUCCUACGUGCGGAUUUGCUAACAACCCGCACGGGCCCCACGUGAUGGUGUAUCGGCCGCCGUCUCACGGUGGAUGUAGUUCUUACAUUGUGCUCGGAAUAGAAAACAUUAAUCACCUAAACUUUGUAUGCAGCACCGCUAACGAGGCGAACGCGGGUGCCUACCACACACUCGGCCCGAAUCGAUCGGUGCUGCUUGCGAUGCCAAGUACAAGAGAAACACCAUUAUGUUUAACGGUGGCGCUUCAAACAUCGCCGAACGGCAACUCCUCACAAACAUGCAUUAACACCACGGAUAGAGAUGAAAGACCUAGCGAAGUCAAAUUCAGUACCGUUCCCUCUCUGUCCGCGCUGCGAUUUUCCCCAUCUCAAAGAACCGGACGGUAUGUCUGUACCAUGUGGAGAUGGGCUGGAUAUGUUUCUGCCUGCCAAGAAAAGUAUGACGCUCUCACGAUUCGGUGCCAAUAUGGCUCAAUGGGUACCCUCCGCCAUCUUAAGGCAUUACAAGAUACUAGUGUCCAAAAAAUGACGCUUUUGGUAUCCUCCCUUGAGACUUUUGUCCAGAAUGCCACAAAAGUUCAAAGUCCGCUCAUGUCAGCGAAUGCUCCAGUCACAACAGCUCUUGCAGUAGCAAGUUUGGAUUAUCCUUUGGCAUGGGUCCCUAAUGCUUGGAGAGCCAUUCGUCAGAUUGCGUUGUGCGUCAAAUUGAAGUGGCAUGUGAUCGCUUGUCAGGAUUUUCUCCGCUGCUGGCAUUCCUAUUUCCUGGGGGCGGACGAUGGCUCUAAGAUGUAUCGCGGAUCAGUAGAAACAAUAAAAAGAGACUCUCGGUCUAAGAGGGUGGUGGGCAGCAAGAAAUCGAUGCGGGCGGUUAUGGCGUGGAGAUGUCGGGUGCGCUGUAGUGUAACUUUGGUCCCUGAGCCGUAUAGGUCGAGAUACUGCCGUCCCUUAUGUUUUAAUGACGUGGUCCACUUAAGGUCCACAGGGAGAUGUGCCGGGGGGCGAUCUUAUCACAAAGCCGCCGGGCGUCGCACCCACGCCGUGGACCUUGAACACACGAUUUCAUCGCAUGGAUCCCUCGUCACGAUGUGCGUCGAUGCUCGUGUAAUCGCUAAUGCGGUCAUUACGCAUAAGUGUGGCCUCUAUACUCCGCUCUCACGCGCCAAUUACGUCUGUAGAACCGAGAGCAGCCUCACCAAGUCCCCAAAGCUUACUCCCACCAACACACAUGUUUCAGAUGCACGCAAUCUUUUCAGCAGACUCAGUUCAUCUCCGUCGUCCGGAGCGCAGCUGAAGGAUUCUGCGUAUUUCCUACGCGGUCCGAGGUCAUUGAGGGACUGCUUCGUAGCUGCGAGGCGUGUAACGACCCAAGACAGCGAUGAGGGGGUCAUGGAGCCUACGAGAUUAAGAGAGGAUCCAGGGGACGUAGAUCUCAGGGAUGUGAGAUUGCACCUGUUACCGGGGUCUGCAGAUGACAAAUUAAUACCUCGCAAGGACUGUACGCGUCGAUGCCGGUGGCCUUCGACGUUUAAAGCCCUGCGACUGCAUUAUACUAGUGUCCGAACGGUUAGAGUUAUAGCACGAGUGGUGUUUUUGAAACCACAGGUUGGGUGCAAUGACUUGUGCGUAUUAUCCAGUUCCCAACCGGCAUCACUUCUAUGCGUAGGCCAUGCGUCACUAGGUCUCUCCUCUAAUGGGGCAAACCGACACUAUUAUUUUCACCCGUUGGGAGGGUCUCCUGAAUGUUGGUAUCCGUUCGAUUCAGUGUCUUACAACCUUCUGAGCCAGGGCUCAGGGCAUGUCCCGUCACAAAAUUUUACGGAACUAUCUACGUAUUUUGAUGCAAACUCCAAAAGCCCUACUCCACUCGGAAGGAUUCUGGGCGCAAAAUUGGCGUUUAACAAGGCUGAUUUACGAAAUCCCCCAUCUCCAUCCCUGUCACUCUCCUCCUACCGCUUACGUUCUCCCCUACCUAGAUAUAAUGUCCCCAUCCCGGAAUGCACGAACCCAACGCUGUAUGUCGGCAUGGAUCUAAGUUUCGUUUACAAUUACUUGCAACUAGUCUCGACCAAAUUACUUAGUACAAUGCGGCUCGGAGUUAUAAAAAGAGAUGUGGGUAUGAAGACCACACAACCCUCGUGUAGUCAUCUAAGCUCAGUGCACUCGCGAUGUAGGUGGUGGCGAAUUGUCAGUGGCACGAUGGAAUGGUACUAUAACGGGCCGGCACACUACAGCGCCCUCAGGCUAACAAGAGGACUGUGCAACCAGAACUUAACAGACGGUUUACUACCGACCGUUACCGUUGACAGACGUAGUAGUGUGCCAAUUUUGAAAGGAACGGCACUGUUGUCAGACCUAUUCGACGCUCCUUCGCCAGUACAUGCGAUACCAAAGAGUCAGGGAUUGUUAAAACCAAAACGGCGUCUUCAUGAGCAUCUCGCAGAUAGUGGGUACUUACUGGGGCGGAUAGUGUUAAGGUAUUACUCGCUUGAGGACCGUGACCCCGUUCAUACGCGAAAACAGGAAUGUUCUACACUGGUGUGGCCGGAUGCCAACUUAUCCUUUGAUCGCGCGUAUAUUGUUCGUAUCCAUUUCCCGGCUCGCCAGUCAGAUUUAUCCAGUAGCUGGCUGAACGGGAGCAUUCGCGCACUUACCAGCCGUCCCCGCCGCGGGCGAAGCGCCUUACACCUUGUGAACUGUCGCCUACCACUCACAUACAAGCCCAGACCCCCAGCAUAUGGAGAGGGAUUGUGCAUCAUGGCUACGCCCUUACUACGCCAACGCUCUAUACCGAUACCCUCCCCUUUUUGCCUCGGGGAAGCCACGCUUGAACAUGGCCCUCUUGAAACGACGUGCCCUUACGUUCUACUAGUACUCAACGCUCGAGAUUAUGUAUUCCUGCGGCUGCCACGAAGAGGUUAUCCUAUUACCUUAGUUGGAGAAUCACGUUCCUAUAAGGAAGCACUAAAAUUGGGGGAAUCUGUGUAUCACAUGGAUUGCAUCGGUCAUACAACUGUUAUCAAAGCGCGCCAGACGUGUGCAGUCCUUGCACUCCCCAACGUGCGCUGCUGCUGCCCAGUGAGUGGACUUAUCUUUUACAUUUACACGUCUGCUUCGACCCCCCAAGCCGACUAUUUGCUUCUGAUCUACGUGCAAGUACCGGCUCUGAAUAUGAGUUAUGGUCACUAUUUGCCAGGGUGCAGUGCAAUCCGAAGUCAACGCCACUCUAGCACUUGGGUUGACCUGUUGGCUACAUAUGCCAUGUCUGUGGACGAGUUUUGGCCUGUAGUGAAAGAGCCUGUAUGGCCUCCACUAGAUGUAGAUCUGCAAAUAGCGCGAGACCAAGAAUCGCAAACUCAGCAACCGGUAACGUUGGCACUUUUGGUGCUUCUCCGGCGGGGUGAACCUAUACUAUCGUUUUGGUGGGAUGUUUCCAGUGCAAUUCAGGGUUCUACCUUUCUCCGGCGACAAACCGAUGUGUUCUACUAUACUCGAGGGCGAGGCGGGCCACGGGGACGGAGUCUGAUGCGCUCACGCAUUGUUGACCCGACGUCCGCUUUUUACCACACGUUAGGCGUACGCAACAAGCGUCUCGCAACUUGUACGUCAAGCGACGGAUGGCAGGGGGGCUACCGUAUCACCCGCCGUCUGUAUGGAACUCGCAUGGGUACUAUCCACCUGGGUAGUCCAAGACAUAUUAACAGUCAUUUCGCGCAUCGAUGUGGAGCUUUUCGGUUCUCGGUUACGAACGAAGACGUGCCAAGCAUCACACCCUUGUAUCAGUUUAUGAAACCCGAGGUCUCUUUCGUCCAAAAUAGUACCCUCAAACUAUCUAAUACCGAUCCCACCACCGUCGCCAUAAGAGGGUCGCUGGGUAUGCCCAAUCAUUACAGUAGCCGGGAUGCGGACGUACCGGCAAUGCCCGUCUUGUUGGCAGCUCCAGGCCGGAUCGAAUUGCCAUCUGUUAUCAGGAGCCCUCUAUUUAUACGUUCAUAUCCGCCCGGAAGCGAAGCCGACCAACCACACCCUCCUUUCAUCUGUCCUUUUUGCGGUCUAAGACACUCCAAUGUUCCGGCACGACAAUUCGUCACAAGUAGUGCAACAGAAACUCCUGGCUCCAAACACGAUGCGCGUCUACUACCUGAAUACGUAAUCGAUACCUAUCGUAUACGCCGUGAGGAUCAGAUGAAUCUUCCUCCACGUCCCGCUCAGGCUACGGACGAGAGCGCGAUACCAAAUGCAUUUGGGACUUCUUAUAUGUCGAUGCCGAGUUGUCUUGGGUGUAAGGUUGUAAACACGGAGUGUGCGGCGAAGGGGCAGCUGGUCUGGCCAAUCGGUACGGUGAACACGCCCUAUCGCUACGAUUCUUCAAUCGAUCUACUCGUGGGCGCUCAAGAUGAACCACGUCUCAAGACUUGUGCACCCCUAGUCCCCUUAAUGCGAAAACCCCUUAGCGAUCCCCUGAAAAUAGAUUUUCGGAGGAUGUGCAGAGGCUCGGAGACGCACGGUAGGAGGUUCCGCAGCGCAUCAUUUGCGCAAGGAUACCGUGCCUGUGAUCGCGGGAAGCGAGAACCCUCUACUAGUAGAUAUCUCAAGGCUGAACGCAUAAGUCUGACAUCGCCCGAGCGUUCACUGAGCGACAAAUACCCGGGGCGACCCAAAGGUUGUAGACGUGGCCGGAUUGGGCAUACAAGAGUUGUGCAACAGAGGUUGGGACGGAGCAUGAUAUGUCGGAUGAAGGGCAAUAGUCGACAGCUUUGGCGAGUAACUUAUUACCGUCGCUGUUUCCUUCUUCAUGCGGGUUCGCCUCUCGACUUCCCAGGAAGGGUAGACUCGCCCGUGCCGUUAGGUCGGGUAAAACCCCCAUGCAGAGCCAUGCCCUCCGCGGCAACUCAAGCCAUAAUGGGGGUACUAGUUAGGUGCCGAGGGGGGCAAUUUCGCUGGCACGGGAACGUGAUUGAUCUAAGUCAACUAUCUGUAAGACGGCUAAAUGGCUUGCCUCACUUGUUAGUACACACAAGGGCAGCGUACUAUACGCUUCAGCGGUUAGGGCAACGAUCAGCCCUCCGAUCUCGAUGCCCCUCUGGGAUUGGGAAACCAGAUCAACGGAGUCGUCGAGUGUAUACACGGCCAUAUUUUGUUUCAUUGCUCGUCCCAACGGUAACUCGAAGCUCCUUAGGAUACUUCCUAGCCCUUGACAUAUGGACCAAACGUCUGGUCCUCACGGUUCCGCACCGCCCGAACACACGACGGGUACAAUUUGGCGAUUUUGCCUCUCACUCUGGCCUUGCUUCCUUACCACAACCUGCACAUGGGUUAGAGACCGGAGUUCAAACUAGAUAUACAACCAUCUGCUCUGCCAUGGUGGUGCAAACGGGUUACAGAGCGCGCGGUUAUACAUCGUUUGCUUUAGACAAACACCGGCUAGAAUUCCCAGAUGUCGCAAAUACGGUUAUCGGCGACAUGACACCAGAACAUCAAGAUCCGGAGCAUCAUCGAUACGUUAAUCAAACAUGCCGUUUCAUAACUUGGGCUAGCGGGUCAUGUAUACCCCCGCAAACGCGGUCUAAUUACCCGGCUCUCUCGUGUGCACAUCAGCACACAGGCAUUGUUGUGUCAACUAACUGUGAUCAGGAUAGCUUGAAGCGCCAACGCUCCGCGUUAUACACUAAGCCAUAUUCCAGUAUGAUUCGGACACAAAAUGACUUGGUGGGGUUUCAUUUCCGGACAUGUGGGAGAGAGCAACGCUUUAGUACCCAGCUAGAUGGUCACGUUAUGAACUUAUCGCAAGUGAAAAGCGCGCUGACAAGGAUCUCCGCUCAGUUGGGUGGGGUGUUAAUGCCCUGCAAACCUAAUUCAGUCUGCUACGGCUCGGAUCUCUCUGUAUACAUUUUCGCUGCAGGCCUCAUAGUCAGGAUGGAAGCCAAGCACACCACUCGCAAUGCCGCGAUAGCAAGUGCCCCAAGAUAUAGAGGACCAUACUUAAUAAGUCCUGGCGUCCUCCGGUUCGGUCGCCGGCGUGCCGUAUCCUCACACCGCUCUAUGAACCUGUAUUUAUCGACUCGACACGCCAUGUUGGGAGUGACUUGUCUGAGCGGCAAAAGACAUAGUGGUGUGAUGACUCUUAGAUGGUUUGGUCACAUGCGUGCCUCACACGUGACAUUCAGUUCAGCGCAACGACAGGCACUUAGAUGUACGCGCCUACGCACACAUGCCCUUAAUCCGGUUACCGCAUGUGGGGUUCAGUACAUAGCUCAUGCGAGGUAUCGAGCGUGCGACGGGAAUAGGCCACGGAUUUCUGAACAGGGUGACUUAGGUGCUACCCCGCUCUCGUUGAUCCGCAUCACAGAUCUUUUUGACACAGGGUAUUAUAUUCUUAACCCGAGCCUGGCAUGUGACGGAGCACAUAGAAUGGAACCUCCGCACAACUUUAAAACGUACAAUGGGCAAACAUCGGUCCUAUGUGAGAAUUUAGAUGUCCGUUAUCUCGUGACGAAAUCUUUUUACAGAGUCGGUCUCCUAUAUCCUGGUCGGAAACGUCGGGCUGUAAUUUAUCCACUGGCUCCUCCUGCCGCGUUAUUAUGCCAUGACCAAGCAACCGCAGAGGUAUGCUCUGCCUAUCAGAGCGUGACCACAAGUAGGCCAGCAAGCCGCCGCGCGUAUAAUUCGCCGACGUCCGAGGAUGGGCGCAUCUUUCCGAGUGUCCCAGUCAGAGUGACAACUUCUACGGCGUUCUUAGAAGGAACGAUAUUUCAUCCGCAUUUUGCAACUGUCUGGGGCUGUUGGCUGGGGAUAUCGCUACCGAUGAUUACAGCUGGAAGUCAUGCGCUACUCUUCGAGUGGGAGUCCAUCGCCGGGCUGUUCUUGGUCGUGCAUACCUCGACUCGUCGGACGAGUGCUGGUCACUUUUACUCGCGCCGUGCAGAUUAUCUUGGGACCGGUUGGGGGCUCGUCCAGCUAAUGGAAACCGGUCACCCCGUGUUCUUUGUCAAUGCAAGGGUUAUGCGUGUUAUUACCGCUACUGCACAACCCACGCGUCAAGGUAAUACUGAAGCGUGGACUUUGUCCAGUGCGCACGGUUCUGAUUUCACCCCCAAGUGGAGGCAUCCGUGGGGAGCCUAUAACACGGCUUUUCCCCGCGGUUCUCUCCUGCUCGAGGAGCGAUCCUUCGAGGUGGGGUGUAAGGUAAACAUGUCAUACGUCGGUACCAAUCCAAGAGCCCACAACCUUGGAGCUAGGCGGCAGAUCUGCAUGAGAAAUCGCGGUUUGAUUGAAAACAUACGAUACUGCGAUGAACUUUCCGCUCAUAGUCGCAUGCAAAGCAACAUCUUCUUAGGUUGGUCCAAGAGUCCAAGCGAAUACAACAGUUAUCCCUGUAGACGCUUAUUGGGGCAACAAUGGUUGGGAAUAUGUGGUAGGACUAUGGAAAAAGCAAAUGCGGAGAACGGCGCUACCUAUGGGCCCCUUUUUGUCGUGAGCCAAACCGGCGUGAUUUCAUCCCCUGAGACUCUCUAUAAAAGUAGAGAAACCCGUUACUAUUUGGGAACUGUUCAACACCUGGGGUGCCUAGGUUGUGCAGUUCGUAUAAAGCGCCGAGGGCAUUCUCGCACGAUUCGCCGGCUACCUGGCUGCCGCCUUCGAGUUCGCCCAGUAUCCGUGGCGGACUCCGCCGUCGAACGGACGGGGGAGAGAGCCCUAUGUACCGAGCCCACGAUCAACAAGGUACACAAAGCGUCUUCUUCUGUCACUAGCUGUGUGGGGACGCCCGGUCGUGGUUAUUCGUCCGAAGGUCGGCGUAGAGUUCAAGCCGACAUACGUGGUCGAGCCCUCCCCACAUGGUACCAGUCAACUACGGAACGGUCAGCCAAAUCGCUGGUGAUCGCGGUAUGUAAGUCCCCUACUUCCACUACUAAGCUAUCGAAAUGCAAACCUCCGAUACAUUCUAGAUGCCUAGUGCCCUUUUCUCUGUAUUUUAAUUUAGUACUAAAACUUAUUGGUAAUUCUAAGAUUAGGAUUCUUGCGCCGGAUAACCAUCCCCCUGUACUAGCGGCUAGAGCGGGUUUGAACUUAACACGAAGCCUUGGCGAGUAUGUCACACAACCAAAACACCCUACCGCAGUAAGCGUUCGGUGUUACUCGAACAUUUCAGCGUGCCUCUGGACAGCUCCGUGUGCCUUGUCCCCGUGGUAUGACAGGCUUGUGUUCGAAACUUUACUGGCGUACAUAUUCCGUUGCAUUGCAUCAGUACUUACAUUAAUCACAAACGCUCUUAUGUUAGCCUCGGAGCGUGCUUAUACCAAUACCAUUGGUAGCCAACGGUCUCUGGGGGGUUGCGCAAUCGAGGGACCUUCUACAUGUGCCUAUUGCAAGGAACGUAUGAACGACAUACCCACGAAGAGCUACCAGUCAAAUUGCUUGACAUCAGUUAUGGGGCGAAAAAGCACACAUUGGUUUAAAAGCAGGGGCUACAAUUGUUUGAGCUUACAUUACAAGACCGAUGUUGACCGGUGGUUUGGCUGUUCGUUAGUUGAGGACUUAUCGCAGUACUGUAGCGCAGACGUCGAUUUCCAGUCCCAGAUUGAGGCGCUCCCUCGUUUAGUAUUACGGGGGGCUACUCUUCCAAUGUUCCCCAAAUCAACACGGUAUACCGUAACUAAAAACACGGCCGGACACCUAGGUGAAUCCGACUACGAGGCGCACGCCACCGCGGGCCAUGUCGUAGGAUCAAGGGUUCGACCCAGACUCUUCCGAUUAGCCUACUCAUACGGACAUCAUAGUAGAUGGGUGUUAUUGGGUCCGCGAGACGGUCCCGUUUCAGGACCAGAGAUGACGACCACAGUGGUACUUUGUUCGCACGACUUCACGGUCAGAGACCCAAGUCGAAGAAUUAGUCCGGCUUGCUCCCGGGCUCACAAGUAUCUAUAUCUUCGGAUGACGCGACCGCAACGGCUCGCUAUGAGGAAUGAAGCUGGGUCACGGCCUAUCAUAGAGAACGGGUGUGCAGCCCAUUGUUGUGUUAAGUUGGUGAAUCUUCUCCGAUGUGCGCGACGUGAACCGCGUACGAAUCAGUUUAGCAACGGGGGGGAUAGACUCGAAGACUGGAGGUGGGGGCAUGUACCGUGUCGAACAACCUGGUCGACGGCUCGGUGUGGUCUUCACACGGAACACCCUGUUACGUCGGCCUUAUGCACUUUUUCCGUUCACGCCAAAUAUAAACUUAGUUCAUGUUCCCUUGCGGGAUAG"))))
}
