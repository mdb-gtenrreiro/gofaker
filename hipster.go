package gofaker

import "errors"

// HipsterWord will return a single hipster word
func HipsterWord() string {
	return getRandValue([]string{"hipster", "word"})
}

// HipsterSentence will generate a random sentence
func HipsterSentence(wordCount int) string {
	return sentence(wordCount, HipsterWord)
}

// HipsterParagraph will generate a random paragraphGenerator
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGenerator(paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, HipsterSentence)
}

func addHipsterLookup() {
	AddFuncLookup("hipsterword", Info{
		Display:     "Hipster Word",
		Category:    "hipster",
		Description: "Random hipster word",
		Example:     "microdosing",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HipsterWord(), nil
		},
	})

	AddFuncLookup("hipstersentence", Info{
		Display:     "Hipster Sentence",
		Category:    "hipster",
		Description: "Random hipster sentence",
		Example:     "Microdosing roof chia echo pickled.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			return HipsterSentence(wordCount), nil
		},
	})

	AddFuncLookup("hipsterparagraph", Info{
		Display:     "Hipster Paragraph",
		Category:    "hipster",
		Description: "Random hipster paragraph",
		Example:     "Microdosing roof chia echo pickled meditation cold-pressed raw denim fingerstache normcore sriracha pork belly. Wolf try-hard pop-up blog tilde hashtag health butcher waistcoat paleo portland vinegar. Microdosing sartorial blue bottle slow-carb freegan five dollar toast you probably haven't heard of them asymmetrical chia farm-to-table narwhal banjo. Gluten-free blog authentic literally synth vinyl meh ethical health fixie banh mi Yuccie. Try-hard drinking squid seitan cray VHS echo chillwave hammock kombucha food truck sustainable.<br />Pug bushwick hella tote bag cliche direct trade waistcoat yr waistcoat knausgaard pour-over master. Pitchfork jean shorts franzen flexitarian distillery hella meggings austin knausgaard crucifix wolf heirloom. Crucifix food truck you probably haven't heard of them trust fund fixie gentrify pitchfork stumptown mlkshk umami chambray blue bottle. 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan hoodie +1 blue bottle. Fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk.<br />Shabby chic typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it. Raw denim cliche dreamcatcher pug fixie park trust fund migas fingerstache sriracha +1 mustache. Tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid stumptown organic schlitz. Flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy. Chia mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Display: "Paragraph Count", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Display: "Sentence Count", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseparator", Display: "Paragraph Separator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			paragraphCount, err := info.GetInt(m, "paragraphcount")
			if err != nil {
				return nil, err
			}
			if paragraphCount <= 0 || paragraphCount > 20 {
				return nil, errors.New("Invalid paragraph count, must be greater than 0, less than 20")
			}

			sentenceCount, err := info.GetInt(m, "sentencecount")
			if err != nil {
				return nil, err
			}
			if sentenceCount <= 0 || sentenceCount > 20 {
				return nil, errors.New("Invalid sentence count, must be greater than 0, less than 20")
			}

			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			paragraphSeparator, err := info.GetString(m, "paragraphseparator")
			if err != nil {
				return nil, err
			}

			return HipsterParagraph(paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})
}
