type MagicDictionary struct {
	dictionary *[]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dictionary = &dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range *this.dictionary {
		var count int = 0
		if len(word) != len(searchWord) {
			continue
		}
		var i = 0
		for {
			if i == len(word) {
				break
			}
			if word[i] != searchWord[i] {
				count = count + 1
			}
			i = i + 1
		}
		if count == 1 {
			return true
		}
	}
	return false
}