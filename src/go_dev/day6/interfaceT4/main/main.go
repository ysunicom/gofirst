refix = "byte "
		case STOSB, MOVSB, CMPSB, LODSB, SCASB:
			prefix = "byte "
		case STOSW, MOVSW, CMPSW, LODSW, SCASW:
			prefix = "word "
		case STOSD, MOVSD, CMPSD, LODSD, SCASD:
			prefix = "dword "
		case STOSQ, MOVSQ, CMPSQ, LODSQ, SCASQ:
			prefix = "qword "
		case LAR:
			prefix = "word "
		case BOUND:
			if inst.Mode == 32 {
				prefix = "qword "
			} else {
				prefix = "dword "
			}
		case PREFETCHW, PREFETCHNTA, PREFETCHT0, PREFETCHT1, PREFETCHT2, CLFLUSH:
			prefix = "zmmword "
		}
		switch inst.Op {
		case MOVSB, MOVSW, MOVSD, MOVSQ, CMPSB, CMPSW, CMPSD, CMPSQ,                                                                                                                                                                                                                                                                                                                                      