package langsetter

import (
	"gopkg.in/ini.v1"
)

// LangSetter sets the text for certain language
type LangSetter struct {
	fileName string
	cfg      *ini.File
}

// NewLangSetter creates a new langsetter
func NewLangSetter(
	fileName string,
) *LangSetter {
	return &LangSetter{
		fileName: fileName,
		cfg:      ini.Empty(),
	}
}

// SetText sets the text for certain tag for certain language
func (ls *LangSetter) SetText(
	lang string,
	tag string,
	text string,
) error {
	langSection := ls.cfg.Section(lang)
	langSection.NewKey(tag, text)
	err := ls.cfg.SaveTo(
		ls.fileName,
	)
	return err
}
