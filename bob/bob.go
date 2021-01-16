/*
Package bob contains one function Hey that describes the responses bob
replies with when certain phrases are said to him.
*/
package bob

import "strings"

/*
Bob answers 'Sure.' if you ask him a question, such as "How are you?".
He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying
anything.
He answers 'Whatever.' to anything else.
*/
func Hey(remark string) string {
	isQuestion := strings.HasSuffix(remark, "?")
	isExclamation := strings.HasSuffix(remark, "!")
	isAllCapitals := strings.Compare(strings.ToUpper(remark), remark) == 0

	if isAllCapitals && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isAllCapitals && isExclamation {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
