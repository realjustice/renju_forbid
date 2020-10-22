package ren

type line struct {
	x []int
}

func newLine(a []int) *line {
	line := new(line)
	line.x = make([]int, S+4)
	line.x[0] = 1024
	line.x[1] = 1024

	for i := 0; i < S; i++ {
		line.x[i+2] = a[i]
	}
	line.x[S+2] = 1024
	line.x[S+3] = 1024

	return line
}

func (line *line) A5(p int) int {
	p += 2
	x0 := line.x[p]
	if x0 == 0 {
		return 0
	}

	xMin := 0
	if 2 < p-4 {
		xMin = p - 4
	} else {
		xMin = 2
	}

	xMax := 0
	if S-3 < p {
		xMax = S - 3
	} else {
		xMax = p
	}

	for i := xMin; i <= xMax; i++ {
		num := line.x[i] + line.x[i+1] + line.x[i+2] + line.x[i+3] + line.x[i+4]
		if num == 5*x0 {
			if !(x0 == 1) {
				return 1
			} else {
				if line.x[i-1] != x0 && line.x[i+5] != x0 {
					return 1
				}
			}
		}
	}
	return 0
}

func (line *line) B4(p int) int {
	p += 2
	x0 := line.x[p]
	if x0 == 0 {
		return 0
	}
	xMin := 0
	if 2 < p-4 {
		xMin = p - 4
	} else {
		xMin = 2
	}

	xMax := 0
	if S-3 < p {
		xMax = S - 3
	} else {
		xMax = p
	}

	for i := xMin; i <= xMax; i++ {
		num := line.x[i] + line.x[i+1] + line.x[i+2] + line.x[i+3] + line.x[i+4]
		if num == 4*x0 {
			shape := (line.x[i] << 4) + (line.x[i+1] << 3) + (line.x[i+2] << 2) + (line.x[i+3] << 1) + line.x[i+4]
			if !(x0 == 1) {
				shape = -shape
				switch shape {
				case 0x1e: //OOOO_
					return 1
				//return COMB(1,i+4);
				case 0x1d: //OOO_O
					if i <= S-7 {
						if line.x[i+5] == 0 && line.x[i+6] == -1 && line.x[i+7] == -1 && line.x[i+8] == -1 {
							if p == i+4 {
								return 2
							}
							//return COMC(1,i+3,i+5);
						}
					}
					return 1
				//return COMB(1,i+3);
				case 0x1b: //OO_OO
					if i <= S-6 {
						if line.x[i+5] == 0 && line.x[i+6] == -1 && line.x[i+7] == -1 {
							if p == i+4 || p == i+3 {
								return 2
							}
							//return COMC(1,i+2,i+5);
						}
					}
					return 1
				//return COMB(1,i+2);
				case 0x17: //O_OOO
					if i <= S-5 {
						if line.x[i+5] == 0 && line.x[i+6] == -1 {
							if p == i+4 || p == i+3 || p == i+2 {
								return 2
							}
							//return COMC(1,i+1,i+5);
						}
					}
					return 1
				//return COMB(1,i+1);
				case 0xf: //_OOOO
					return 1
					//return COMB(1,i);
				}
			} else {
				switch shape {
				case 0x1e, -0x1e:
					if line.x[i-1] != x0 && line.x[i+5] != x0 {
						return 1
						//return COMB(1,i+4);
					}
					break
				case 0x1d, -0x1d:
					if line.x[i-1] != x0 && line.x[i+5] != x0 {
						if line.x[i+5] == 0 {
							if i <= 8 {
								if line.x[i+6] == x0 && line.x[i+7] == x0 && line.x[i+8] == x0 && line.x[i+9] != x0 {
									if p == i+4 {
										return 2
									}
									//return COMC(1,i+3,i+5);
								}
							}
						}
						return 1
						//return COMB(1,i+3);
					}
					break
				case 0x1b, -0x1b:
					if line.x[i-1] != x0 && line.x[i+5] != x0 {
						if line.x[i+5] == 0 {
							if i <= 9 {
								if line.x[i+6] == x0 && line.x[i+7] == x0 && line.x[i+8] != x0 {
									if p == i+4 || p == i+3 {
										return 2
									}
									//return COMC(1,i+2,i+5);
								}
							}
						}
						return 1
					}
					break
				case 0x17, -0x17:
					if line.x[i-1] != x0 && line.x[i+5] != x0 {
						if line.x[i+5] == 0 {
							if i <= 10 {
								if line.x[i+6] == x0 && line.x[i+7] != x0 {
									if p == i+4 || p == i+3 || p == i+2 {
										return 2
									}
									//return COMC(1,i+1,i+5);
								}
							}
						}
						return 1
						//return COMB(1,i+1);
					}
					break
				case 0xf, -0xf:
					if line.x[i-1] != x0 && line.x[i+5] != x0 {
						return 1
						//return COMB(1,i);
					}
					break
				}
			}
		}
	}
	return 0
}

func (line *line) A3(p int) int {
	p += 2
	x0 := line.x[p]
	if x0 == 0 {
		return 0
	} else {
		xMin := 0
		if 2 < p-3 {
			xMin = p - 3
		} else {
			xMin = 2
		}
		xMax := 0
		if S-2 < p {
			xMax = S - 2
		} else {
			xMax = p
		}

		for i := xMin; i <= xMax; i++ {
			num1 := line.x[i] + line.x[i+1] + line.x[i+2] + line.x[i+3]
			num2 := line.x[i] * line.x[i+1] * line.x[i+2] * line.x[i+3]
			if num1 == 3*x0 && num2 == 0 {
				shape := (line.x[i] << 3) + (line.x[i+1] << 2) + (line.x[i+2] << 1) + line.x[i+3]
				if !(x0 == 1) {
					{
						shape = -shape
						switch shape {
						case 0xe: //OOO_
							if line.x[i-1] == 0 && line.x[i-2] != -1 && line.x[i+4] != -1 {
								if (line.x[i-2] == 0) && (line.x[i+4] == 0) {
									return COMC(1, i-1, i+3)
								}
								if line.x[i-2] == 0 {
									return COMB(1, i-1)
								}
								if line.x[i+4] == 0 {
									return COMB(1, i+3)
								}
							}
						case 0xd: //OO_O
							if line.x[i-1] == 0 && line.x[i+4] == 0 {
								return COMB(1, i+2)
							}
						case 0xb: //O_OO
							if line.x[i-1] == 0 && line.x[i+4] == 0 {
								return COMB(1, i+1)
							}
						}
					}
				} else {
					switch shape {
					case 0xe, -0xe: //XXX_
						if line.x[i-1] == 0 && line.x[i-2] != x0 && line.x[i+4] != x0 {
							if (line.x[i-2] == 0 && line.x[i-3] != x0) && (line.x[i+4] == 0 && line.x[i+5] != x0) {
								return COMC(1, i-1, i+3)
							}
							if line.x[i-2] == 0 && line.x[i-3] != x0 {
								return COMB(1, i-1)
							}
							if line.x[i+4] == 0 && line.x[i+5] != x0 {
								return COMB(1, i+3)
							}
						}
					case 0xd, -0xd: //XX_X
						if line.x[i-1] == 0 && line.x[i+4] == 0 && line.x[i-2] != x0 && line.x[i+5] != x0 {
							return COMB(1, i+2)
						}
					case 0xb, -0xb:
						if line.x[i-1] == 0 && line.x[i+4] == 0 && line.x[i-2] != x0 && line.x[i+5] != x0 {
							return COMB(1, i+1)
						}
					}
				}
			}
		}
	}
	return 0
}

func (line *line) A6(p int) int {
	p += 2
	if line.x[p] != 1 {
		return 0
	}
	xMin := 0
	if 2 < p-5 {
		xMin = p - 5
	} else {
		xMin = 2
	}
	xMax := 0
	if S-4 < p {
		xMax = S - 4
	} else {
		xMax = p
	}
	for i := xMin; i <= xMax; i++ {
		num := line.x[i] + line.x[i+1] + line.x[i+2] + line.x[i+3] + line.x[i+4] + line.x[i+5]
		if num == 6 {
			return 1
		}
	}

	return 0
}

func COMB(x, y int) int {
	return (x << 8) | (y - 2)
}

func COMC(x, y, z int) int {
	return COMB(COMB(x, y), z)
}
