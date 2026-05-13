type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var enc []string
	for _, r := range strs {
		enc = append(enc, strconv.Itoa(len(r)))
	}

	return strings.Join(enc, ",") + "#" + strings.Join(strs, "")
}
// 5,5#helloworld

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}

	parts := strings.SplitN(encoded, "#", 2) // {5,5}{helloworld}
	sizes := strings.Split(parts[0], ",") // {5 5}
	var res []string
	i := 0
	for _ ,sz := range sizes {
		if sz == "" {
			continue
		}
		leng, _ := strconv.Atoi(sz)
		res = append(res, parts[1][i:i+leng])
		i += leng
	}

	return res
}
