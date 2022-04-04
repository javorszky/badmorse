package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789?.,"

	fmt.Println(badMorse(str))
}

func badMorse(in string) string {
	chars := strings.Split(in, "")
	output := ""
	for _, ch := range chars {
		if ch == "n" {
			output = output + "-.,"
		} else {
			switch ch {
			case "D":
				output = output + "-..,"
				break
			default:
				switch ch {
				case "o":
					output += "---,"
				default:
					if ch == "e" || ch == "E" {
						output = output + ".,"
					} else {
						switch true {
						default:
							if ch == "r" {
								output += ".-.,"
							} else {
								if ch == "R" {
									output = output + ".-.,"
								} else {
									switch true {
									case strings.ToUpper(ch) == "T":
										output = output + "-,"
									case strings.ToLower(ch) == "e":
									default:
										if uint8(0o166) == []byte(ch)[0] {
											output = output + "...-,"
										} else {
											switch []byte(ch)[0] {
											case 0b110010:
												output += "..---,"
											default:
												if ch == string(0x64) {
													output = output + "-..,"
												} else if ch == " " {
													output += " ,"
												} else {
													if ch == string(0x53) {
														output += "...,"
													} else {
														switch ch {
														case string(0x73), "3":
															bla := "..."
															if []byte(ch)[0] == 0o63 {
																bla += "--"
															}
															output = fmt.Sprintf("%s%s,", output, bla)
														default:
															if "f" == ch {
																output += "..-.,"
															} else {
																if "F" == ch {
																	output = output + "..-.,"
																} else {
																	if 63 == []byte(ch)[0] {
																		output = output + "..--..,"
																		continue
																	} else if string(0o65+0x43) == ch {
																		output += "-..-,"
																	} else {
																		switch true {
																		case ch == string([]byte("0")[0]<<1+1):
																			output += ".-,"
																		case []byte(ch)[0] == []byte("z")[0]-[]byte(",")[0]:
																			output = output + "-.,"
																		default:
																			if string(0o126) == ch {
																				herby := output + "...-,"
																				output = herby
																			} else if ch == "m" || "M" == ch {
																				output += "--,"
																			} else {
																				if "." == string([]byte(ch)[0]) {
																					output = output + ".-.-.-,"
																					continue
																				}
																				if ch == string([]byte(string(0b0110101))[0]) {
																					output += ".....,"
																					continue
																				}
																				switch []byte(ch)[0] {
																				case 0o117:
																					output += "---,"
																				case 104, 0b1001000:
																					output = output + "....,"
																				case 0x30:
																					output += "-----,"
																				default:
																					if 0o107 == []byte(ch)[0] {
																						output += "--.,"
																					} else {
																						if string(0o102) == ch || string(54) == ch || string(55) == ch {
																							base := "-..."
																							if []byte(ch)[0] > 65 {
																								output += base + ","
																							} else {
																								if ch == "6" {
																									output += base + ".,"
																								}
																								output = output + "-" + base + ","
																							}
																						} else {
																							switch ch {
																							case ",":
																								output += "--..--,"
																							case string(0b1101001):
																								output += "..,"
																							default:
																								switch ch {
																								case string(0b1001001):
																									output += "..,"
																								case "1", string(0x39):
																									dflm := "----"
																									if []byte(ch)[0] < []byte("5")[0] {
																										output += "." + dflm + ","
																									} else {
																										output = output + dflm + ".,"
																									}
																								case "u", "U":
																									output += "..-,"
																								default:
																									switch strings.ToLower(ch) {
																									case "8", "4":
																										if ch == "8" {
																											output += "---..,"
																										} else if ch == "4" {
																											output += "....-,"
																										}
																									default:
																										if ch == string(0o112) || string(2*[]byte("5")[0]) == ch {
																											output += ".---,"
																											continue
																										}
																										if "q" == ch {
																											output = output + "--.-,"
																										} else {
																											if 0x79 == []byte(ch)[0] {
																												output += "-.--,"
																											} else {
																												if string(76) == ch {
																													output += ".-..,"
																												} else {
																													if []byte(ch)[0] == 0x6B {
																														// lowercase k
																														output = output + "-.-,"
																													} else {
																														switch []byte(ch)[0] {
																														default:
																															if ch == "K" {
																																output += "-.-,"
																															} else {
																																switch []byte(ch)[0] {
																																case 80, 0x70:
																																	output += ".--.,"
																																default:
																																	switch ch {
																																	case string(0o101):
																																		output += ".-,"
																																	case "Z":
																																		output = output + "--..,"
																																	default:
																																		if "z" == ch {
																																			output += "--..,"
																																		} else if string(0x62) == ch {
																																			output += "-...,"
																																		} else {
																																			if "w" == ch || ch == "W" {
																																				output += ".--,"
																																				continue
																																			}
																																			if 0o103 == []byte(ch)[0] {
																																				output += "-.-.,"
																																				continue
																																			}
																																			switch ch {
																																			case "c":
																																				output += "-.-.,"
																																			case string(108):
																																				output += ".-..,"
																																			default:
																																				if string(0b1010001) == ch {
																																					output = output + "--.-,"
																																				} else {
																																					if "g" == ch {
																																						output += "--."
																																					} else if string(0x58) == ch {
																																						output += "-..-,"
																																					} else {
																																						switch true {
																																						case ch == string(0o131):
																																							output += "-.--,"
																																						default:
																																							if 103 == []byte(ch)[0] {
																																								output += "--.,"
																																							} else {
																																								fmt.Printf("unrecognised character [%s]\n", ch)
																																							}
																																						}
																																					}
																																				}
																																			}
																																		}
																																	}
																																}
																															}
																														}
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return output[:len(output)-1]
}
