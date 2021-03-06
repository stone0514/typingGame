package internal

import (
	"unicode/utf8"

	"github.com/stone0514/typingGame/pkg"
)

// GenerationInputPattern Document
/*
 *------------------------------------------------------------------
 * @brief GenerationInputPattern
 *
 * @param[in]     []string             target
 *
 * @return[want]  map[string][]string  concatLIst
 *
 * @rules
 * after reading target ([]string)
 *
 *------------------------------------------------------------------
 */
func GenerationInputPattern(target []string) map[string][]string {

	var (
		s, ns          string
		tmpList        []string
		isValidSingleN bool
	)

	concatList := map[string][]string{}

	tmpTarget := target
	d := pkg.Dictionary()

	for i := 0; i < len(tmpTarget); i++ {
		s = tmpTarget[i]
		if i+1 < len(tmpTarget) {
			ns = tmpTarget[i+1]
		} else {
			ns = ""
		}

		if s == "ん" {
			nList := d[s]

			//at the end of character "n" is not allowed
			if len(tmpTarget)-1 == i {
				isValidSingleN = false
				//when a vowel,"な行" or "や行" comes behind "n" is not allowed
			} else if i+1 < len(tmpTarget) && ns == "あ" || ns == "い" || ns == "う" || ns == "え" || ns == "お" || ns == "な" || ns == "に" || ns == "ぬ" || ns == "ね" || ns == "の" || ns == "や" || ns == "ゆ" || ns == "よ" {
				isValidSingleN = false
				//if other case "n" is allowed
			} else {
				isValidSingleN = true
			}
			for _, t := range nList {
				if !isValidSingleN && t == "n" {
					continue
				}
				tmpList = []string{}
				tmpList = append(tmpList, d[s]...)
				tmpList = append(tmpList, t)
			}
		} else if s == "っ" {
			var (
				ltuList  = d[s]
				nextList = d[ns]
				hsList   []string
			)

			// get the consonant next character
			for _, t := range nextList {
				str := string(pkg.GetRune(t, 0))
				hsList = append(hsList, str)
			}
			ltuTypeList := append(hsList, ltuList...)
			tmpList = ltuTypeList
		} else if utf8.RuneCountInString(s) == 2 && string(pkg.GetRune(s, 0)) != "ん" {
			tmpList = []string{}
			//create "きゃ" pattern
			tmpList = append(tmpList, d[s]...)

			// create "き"+"ゃ" pattern
			var fstList, sndList, retList []string

			fst := string(pkg.GetRune(s, 0))
			snd := string(pkg.GetRune(s, 1))
			fstList = append(fstList, d[fst]...)
			sndList = append(sndList, d[snd]...)

			for _, fstStr := range fstList {
				for _, sndStr := range sndList {
					concat := fstStr + sndStr
					retList = append(retList, concat)
				}
			}
			tmpList = []string{}
			tmpList = append(tmpList, d[s]...)
			tmpList = append(tmpList, retList...)
		} else {
			tmpList = d[s]
		}
		concatList[target[i]] = append(concatList[target[i]], tmpList...)
	}
	return concatList
}
