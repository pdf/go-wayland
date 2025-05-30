// Generated by go-wayland-scanner
// https://github.com/pdf/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/text-input/text-input-unstable-v1.xml
//
// text_input_unstable_v1 Protocol Copyright:
//
// Copyright © 2012, 2013 Intel Corporation
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package text_input

import "github.com/pdf/go-wayland/client"

// TextInput : text input
//
// An object used for text input. Adds support for text input and input
// methods to applications. A text_input object is created from a
// wl_text_input_manager and corresponds typically to a text entry in an
// application.
//
// Requests are used to activate/deactivate the text_input object and set
// state information like surrounding and selected text or the content type.
// The information about entered text is sent to the text_input object via
// the pre-edit and commit events. Using this interface removes the need
// for applications to directly process hardware key events and compose text
// out of them.
//
// Text is generally UTF-8 encoded, indices and lengths are in bytes.
//
// Serials are used to synchronize the state between the text input and
// an input method. New serials are sent by the text input in the
// commit_state request and are used by the input method to indicate
// the known text input state in events like preedit_string, commit_string,
// and keysym. The text input can then ignore events from the input method
// which are based on an outdated state (for example after a reset).
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
type TextInput struct {
	client.BaseProxy
	enterHandler                 TextInputEnterHandlerFunc
	leaveHandler                 TextInputLeaveHandlerFunc
	modifiersMapHandler          TextInputModifiersMapHandlerFunc
	inputPanelStateHandler       TextInputInputPanelStateHandlerFunc
	preeditStringHandler         TextInputPreeditStringHandlerFunc
	preeditStylingHandler        TextInputPreeditStylingHandlerFunc
	preeditCursorHandler         TextInputPreeditCursorHandlerFunc
	commitStringHandler          TextInputCommitStringHandlerFunc
	cursorPositionHandler        TextInputCursorPositionHandlerFunc
	deleteSurroundingTextHandler TextInputDeleteSurroundingTextHandlerFunc
	keysymHandler                TextInputKeysymHandlerFunc
	languageHandler              TextInputLanguageHandlerFunc
	textDirectionHandler         TextInputTextDirectionHandlerFunc
}

// NewTextInput : text input
//
// An object used for text input. Adds support for text input and input
// methods to applications. A text_input object is created from a
// wl_text_input_manager and corresponds typically to a text entry in an
// application.
//
// Requests are used to activate/deactivate the text_input object and set
// state information like surrounding and selected text or the content type.
// The information about entered text is sent to the text_input object via
// the pre-edit and commit events. Using this interface removes the need
// for applications to directly process hardware key events and compose text
// out of them.
//
// Text is generally UTF-8 encoded, indices and lengths are in bytes.
//
// Serials are used to synchronize the state between the text input and
// an input method. New serials are sent by the text input in the
// commit_state request and are used by the input method to indicate
// the known text input state in events like preedit_string, commit_string,
// and keysym. The text input can then ignore events from the input method
// which are based on an outdated state (for example after a reset).
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
func NewTextInput(ctx *client.Context) *TextInput {
	zwpTextInputV1 := &TextInput{}
	ctx.Register(zwpTextInputV1)
	return zwpTextInputV1
}

// Activate : request activation
//
// Requests the text_input object to be activated (typically when the
// text entry gets focus).
//
// The seat argument is a wl_seat which maintains the focus for this
// activation. The surface argument is a wl_surface assigned to the
// text_input object and tracked for focus lost. The enter event
// is emitted on successful activation.
func (i *TextInput) Activate(seat *client.Seat, surface *client.Surface) error {
	const opcode = 0
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], surface.ID())
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Deactivate : request deactivation
//
// Requests the text_input object to be deactivated (typically when the
// text entry lost focus). The seat argument is a wl_seat which was used
// for activation.
func (i *TextInput) Deactivate(seat *client.Seat) error {
	const opcode = 1
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// ShowInputPanel : show input panels
//
// Requests input panels (virtual keyboard) to show.
func (i *TextInput) ShowInputPanel() error {
	const opcode = 2
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// HideInputPanel : hide input panels
//
// Requests input panels (virtual keyboard) to hide.
func (i *TextInput) HideInputPanel() error {
	const opcode = 3
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Reset : reset
//
// Should be called by an editor widget when the input state should be
// reset, for example after the text was changed outside of the normal
// input method flow.
func (i *TextInput) Reset() error {
	const opcode = 4
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetSurroundingText : sets the surrounding text
//
// Sets the plain surrounding text around the input position. Text is
// UTF-8 encoded. Cursor is the byte offset within the
// surrounding text. Anchor is the byte offset of the
// selection anchor within the surrounding text. If there is no selected
// text anchor, then it is the same as cursor.
func (i *TextInput) SetSurroundingText(text string, cursor, anchor uint32) error {
	const opcode = 5
	textLen := client.PaddedLen(len(text) + 1)
	_reqBufLen := 8 + (4 + textLen) + 4 + 4
	_reqBuf := make([]byte, _reqBufLen)
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutString(_reqBuf[l:l+(4+textLen)], text, textLen)
	l += (4 + textLen)
	client.PutUint32(_reqBuf[l:l+4], uint32(cursor))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(anchor))
	err := i.Context().WriteMsg(_reqBuf, nil)
	return err
}

// SetContentType : set content purpose and hint
//
// Sets the content purpose and content hint. While the purpose is the
// basic purpose of an input field, the hint flags allow to modify some
// of the behavior.
//
// When no content type is explicitly set, a normal content purpose with
// default hints (auto completion, auto correction, auto capitalization)
// should be assumed.
func (i *TextInput) SetContentType(hint, purpose uint32) error {
	const opcode = 6
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(hint))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(purpose))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetCursorRectangle :
func (i *TextInput) SetCursorRectangle(x, y, width, height int32) error {
	const opcode = 7
	const _reqBufLen = 8 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(x))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(y))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetPreferredLanguage : sets preferred language
//
// Sets a specific language. This allows for example a virtual keyboard to
// show a language specific layout. The "language" argument is an RFC-3066
// format language tag.
//
// It could be used for example in a word processor to indicate the
// language of the currently edited document or in an instant message
// application which tracks languages of contacts.
func (i *TextInput) SetPreferredLanguage(language string) error {
	const opcode = 8
	languageLen := client.PaddedLen(len(language) + 1)
	_reqBufLen := 8 + (4 + languageLen)
	_reqBuf := make([]byte, _reqBufLen)
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutString(_reqBuf[l:l+(4+languageLen)], language, languageLen)
	err := i.Context().WriteMsg(_reqBuf, nil)
	return err
}

// CommitState :
//
//	serial: used to identify the known state
func (i *TextInput) CommitState(serial uint32) error {
	const opcode = 9
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(serial))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// InvokeAction :
func (i *TextInput) InvokeAction(button, index uint32) error {
	const opcode = 10
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(button))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(index))
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

func (i *TextInput) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

type TextInputContentHint uint32

// TextInputContentHint : content hint
//
// Content hint is a bitmask to allow to modify the behavior of the text
// input.
const (
	// TextInputContentHintNone : no special behaviour
	TextInputContentHintNone TextInputContentHint = 0x0
	// TextInputContentHintDefault : auto completion, correction and capitalization
	TextInputContentHintDefault TextInputContentHint = 0x7
	// TextInputContentHintPassword : hidden and sensitive text
	TextInputContentHintPassword TextInputContentHint = 0xc0
	// TextInputContentHintAutoCompletion : suggest word completions
	TextInputContentHintAutoCompletion TextInputContentHint = 0x1
	// TextInputContentHintAutoCorrection : suggest word corrections
	TextInputContentHintAutoCorrection TextInputContentHint = 0x2
	// TextInputContentHintAutoCapitalization : switch to uppercase letters at the start of a sentence
	TextInputContentHintAutoCapitalization TextInputContentHint = 0x4
	// TextInputContentHintLowercase : prefer lowercase letters
	TextInputContentHintLowercase TextInputContentHint = 0x8
	// TextInputContentHintUppercase : prefer uppercase letters
	TextInputContentHintUppercase TextInputContentHint = 0x10
	// TextInputContentHintTitlecase : prefer casing for titles and headings (can be language dependent)
	TextInputContentHintTitlecase TextInputContentHint = 0x20
	// TextInputContentHintHiddenText : characters should be hidden
	TextInputContentHintHiddenText TextInputContentHint = 0x40
	// TextInputContentHintSensitiveData : typed text should not be stored
	TextInputContentHintSensitiveData TextInputContentHint = 0x80
	// TextInputContentHintLatin : just latin characters should be entered
	TextInputContentHintLatin TextInputContentHint = 0x100
	// TextInputContentHintMultiline : the text input is multiline
	TextInputContentHintMultiline TextInputContentHint = 0x200
)

func (e TextInputContentHint) Name() string {
	switch e {
	case TextInputContentHintNone:
		return "none"
	case TextInputContentHintDefault:
		return "default"
	case TextInputContentHintPassword:
		return "password"
	case TextInputContentHintAutoCompletion:
		return "auto_completion"
	case TextInputContentHintAutoCorrection:
		return "auto_correction"
	case TextInputContentHintAutoCapitalization:
		return "auto_capitalization"
	case TextInputContentHintLowercase:
		return "lowercase"
	case TextInputContentHintUppercase:
		return "uppercase"
	case TextInputContentHintTitlecase:
		return "titlecase"
	case TextInputContentHintHiddenText:
		return "hidden_text"
	case TextInputContentHintSensitiveData:
		return "sensitive_data"
	case TextInputContentHintLatin:
		return "latin"
	case TextInputContentHintMultiline:
		return "multiline"
	default:
		return ""
	}
}

func (e TextInputContentHint) Value() string {
	switch e {
	case TextInputContentHintNone:
		return "0x0"
	case TextInputContentHintDefault:
		return "0x7"
	case TextInputContentHintPassword:
		return "0xc0"
	case TextInputContentHintAutoCompletion:
		return "0x1"
	case TextInputContentHintAutoCorrection:
		return "0x2"
	case TextInputContentHintAutoCapitalization:
		return "0x4"
	case TextInputContentHintLowercase:
		return "0x8"
	case TextInputContentHintUppercase:
		return "0x10"
	case TextInputContentHintTitlecase:
		return "0x20"
	case TextInputContentHintHiddenText:
		return "0x40"
	case TextInputContentHintSensitiveData:
		return "0x80"
	case TextInputContentHintLatin:
		return "0x100"
	case TextInputContentHintMultiline:
		return "0x200"
	default:
		return ""
	}
}

func (e TextInputContentHint) String() string {
	return e.Name() + "=" + e.Value()
}

type TextInputContentPurpose uint32

// TextInputContentPurpose : content purpose
//
// The content purpose allows to specify the primary purpose of a text
// input.
//
// This allows an input method to show special purpose input panels with
// extra characters or to disallow some characters.
const (
	// TextInputContentPurposeNormal : default input, allowing all characters
	TextInputContentPurposeNormal TextInputContentPurpose = 0
	// TextInputContentPurposeAlpha : allow only alphabetic characters
	TextInputContentPurposeAlpha TextInputContentPurpose = 1
	// TextInputContentPurposeDigits : allow only digits
	TextInputContentPurposeDigits TextInputContentPurpose = 2
	// TextInputContentPurposeNumber : input a number (including decimal separator and sign)
	TextInputContentPurposeNumber TextInputContentPurpose = 3
	// TextInputContentPurposePhone : input a phone number
	TextInputContentPurposePhone TextInputContentPurpose = 4
	// TextInputContentPurposeUrl : input an URL
	TextInputContentPurposeUrl TextInputContentPurpose = 5
	// TextInputContentPurposeEmail : input an email address
	TextInputContentPurposeEmail TextInputContentPurpose = 6
	// TextInputContentPurposeName : input a name of a person
	TextInputContentPurposeName TextInputContentPurpose = 7
	// TextInputContentPurposePassword : input a password (combine with password or sensitive_data hint)
	TextInputContentPurposePassword TextInputContentPurpose = 8
	// TextInputContentPurposeDate : input a date
	TextInputContentPurposeDate TextInputContentPurpose = 9
	// TextInputContentPurposeTime : input a time
	TextInputContentPurposeTime TextInputContentPurpose = 10
	// TextInputContentPurposeDatetime : input a date and time
	TextInputContentPurposeDatetime TextInputContentPurpose = 11
	// TextInputContentPurposeTerminal : input for a terminal
	TextInputContentPurposeTerminal TextInputContentPurpose = 12
)

func (e TextInputContentPurpose) Name() string {
	switch e {
	case TextInputContentPurposeNormal:
		return "normal"
	case TextInputContentPurposeAlpha:
		return "alpha"
	case TextInputContentPurposeDigits:
		return "digits"
	case TextInputContentPurposeNumber:
		return "number"
	case TextInputContentPurposePhone:
		return "phone"
	case TextInputContentPurposeUrl:
		return "url"
	case TextInputContentPurposeEmail:
		return "email"
	case TextInputContentPurposeName:
		return "name"
	case TextInputContentPurposePassword:
		return "password"
	case TextInputContentPurposeDate:
		return "date"
	case TextInputContentPurposeTime:
		return "time"
	case TextInputContentPurposeDatetime:
		return "datetime"
	case TextInputContentPurposeTerminal:
		return "terminal"
	default:
		return ""
	}
}

func (e TextInputContentPurpose) Value() string {
	switch e {
	case TextInputContentPurposeNormal:
		return "0"
	case TextInputContentPurposeAlpha:
		return "1"
	case TextInputContentPurposeDigits:
		return "2"
	case TextInputContentPurposeNumber:
		return "3"
	case TextInputContentPurposePhone:
		return "4"
	case TextInputContentPurposeUrl:
		return "5"
	case TextInputContentPurposeEmail:
		return "6"
	case TextInputContentPurposeName:
		return "7"
	case TextInputContentPurposePassword:
		return "8"
	case TextInputContentPurposeDate:
		return "9"
	case TextInputContentPurposeTime:
		return "10"
	case TextInputContentPurposeDatetime:
		return "11"
	case TextInputContentPurposeTerminal:
		return "12"
	default:
		return ""
	}
}

func (e TextInputContentPurpose) String() string {
	return e.Name() + "=" + e.Value()
}

type TextInputPreeditStyle uint32

// TextInputPreeditStyle :
const (
	// TextInputPreeditStyleDefault : default style for composing text
	TextInputPreeditStyleDefault TextInputPreeditStyle = 0
	// TextInputPreeditStyleNone : style should be the same as in non-composing text
	TextInputPreeditStyleNone      TextInputPreeditStyle = 1
	TextInputPreeditStyleActive    TextInputPreeditStyle = 2
	TextInputPreeditStyleInactive  TextInputPreeditStyle = 3
	TextInputPreeditStyleHighlight TextInputPreeditStyle = 4
	TextInputPreeditStyleUnderline TextInputPreeditStyle = 5
	TextInputPreeditStyleSelection TextInputPreeditStyle = 6
	TextInputPreeditStyleIncorrect TextInputPreeditStyle = 7
)

func (e TextInputPreeditStyle) Name() string {
	switch e {
	case TextInputPreeditStyleDefault:
		return "default"
	case TextInputPreeditStyleNone:
		return "none"
	case TextInputPreeditStyleActive:
		return "active"
	case TextInputPreeditStyleInactive:
		return "inactive"
	case TextInputPreeditStyleHighlight:
		return "highlight"
	case TextInputPreeditStyleUnderline:
		return "underline"
	case TextInputPreeditStyleSelection:
		return "selection"
	case TextInputPreeditStyleIncorrect:
		return "incorrect"
	default:
		return ""
	}
}

func (e TextInputPreeditStyle) Value() string {
	switch e {
	case TextInputPreeditStyleDefault:
		return "0"
	case TextInputPreeditStyleNone:
		return "1"
	case TextInputPreeditStyleActive:
		return "2"
	case TextInputPreeditStyleInactive:
		return "3"
	case TextInputPreeditStyleHighlight:
		return "4"
	case TextInputPreeditStyleUnderline:
		return "5"
	case TextInputPreeditStyleSelection:
		return "6"
	case TextInputPreeditStyleIncorrect:
		return "7"
	default:
		return ""
	}
}

func (e TextInputPreeditStyle) String() string {
	return e.Name() + "=" + e.Value()
}

type TextInputTextDirection uint32

// TextInputTextDirection :
const (
	// TextInputTextDirectionAuto : automatic text direction based on text and language
	TextInputTextDirectionAuto TextInputTextDirection = 0
	// TextInputTextDirectionLtr : left-to-right
	TextInputTextDirectionLtr TextInputTextDirection = 1
	// TextInputTextDirectionRtl : right-to-left
	TextInputTextDirectionRtl TextInputTextDirection = 2
)

func (e TextInputTextDirection) Name() string {
	switch e {
	case TextInputTextDirectionAuto:
		return "auto"
	case TextInputTextDirectionLtr:
		return "ltr"
	case TextInputTextDirectionRtl:
		return "rtl"
	default:
		return ""
	}
}

func (e TextInputTextDirection) Value() string {
	switch e {
	case TextInputTextDirectionAuto:
		return "0"
	case TextInputTextDirectionLtr:
		return "1"
	case TextInputTextDirectionRtl:
		return "2"
	default:
		return ""
	}
}

func (e TextInputTextDirection) String() string {
	return e.Name() + "=" + e.Value()
}

// TextInputEnterEvent : enter event
//
// Notify the text_input object when it received focus. Typically in
// response to an activate request.
type TextInputEnterEvent struct {
	Surface *client.Surface
}
type TextInputEnterHandlerFunc func(TextInputEnterEvent)

// SetEnterHandler : sets handler for TextInputEnterEvent
func (i *TextInput) SetEnterHandler(f TextInputEnterHandlerFunc) {
	i.enterHandler = f
}

// TextInputLeaveEvent : leave event
//
// Notify the text_input object when it lost focus. Either in response
// to a deactivate request or when the assigned surface lost focus or was
// destroyed.
type TextInputLeaveEvent struct{}
type TextInputLeaveHandlerFunc func(TextInputLeaveEvent)

// SetLeaveHandler : sets handler for TextInputLeaveEvent
func (i *TextInput) SetLeaveHandler(f TextInputLeaveHandlerFunc) {
	i.leaveHandler = f
}

// TextInputModifiersMapEvent : modifiers map
//
// Transfer an array of 0-terminated modifier names. The position in
// the array is the index of the modifier as used in the modifiers
// bitmask in the keysym event.
type TextInputModifiersMapEvent struct {
	Map []byte
}
type TextInputModifiersMapHandlerFunc func(TextInputModifiersMapEvent)

// SetModifiersMapHandler : sets handler for TextInputModifiersMapEvent
func (i *TextInput) SetModifiersMapHandler(f TextInputModifiersMapHandlerFunc) {
	i.modifiersMapHandler = f
}

// TextInputInputPanelStateEvent : state of the input panel
//
// Notify when the visibility state of the input panel changed.
type TextInputInputPanelStateEvent struct {
	State uint32
}
type TextInputInputPanelStateHandlerFunc func(TextInputInputPanelStateEvent)

// SetInputPanelStateHandler : sets handler for TextInputInputPanelStateEvent
func (i *TextInput) SetInputPanelStateHandler(f TextInputInputPanelStateHandlerFunc) {
	i.inputPanelStateHandler = f
}

// TextInputPreeditStringEvent : pre-edit
//
// Notify when a new composing text (pre-edit) should be set around the
// current cursor position. Any previously set composing text should
// be removed.
//
// The commit text can be used to replace the preedit text on reset
// (for example on unfocus).
//
// The text input should also handle all preedit_style and preedit_cursor
// events occurring directly before preedit_string.
type TextInputPreeditStringEvent struct {
	Serial uint32
	Text   string
	Commit string
}
type TextInputPreeditStringHandlerFunc func(TextInputPreeditStringEvent)

// SetPreeditStringHandler : sets handler for TextInputPreeditStringEvent
func (i *TextInput) SetPreeditStringHandler(f TextInputPreeditStringHandlerFunc) {
	i.preeditStringHandler = f
}

// TextInputPreeditStylingEvent : pre-edit styling
//
// Sets styling information on composing text. The style is applied for
// length bytes from index relative to the beginning of the composing
// text (as byte offset). Multiple styles can
// be applied to a composing text by sending multiple preedit_styling
// events.
//
// This event is handled as part of a following preedit_string event.
type TextInputPreeditStylingEvent struct {
	Index  uint32
	Length uint32
	Style  uint32
}
type TextInputPreeditStylingHandlerFunc func(TextInputPreeditStylingEvent)

// SetPreeditStylingHandler : sets handler for TextInputPreeditStylingEvent
func (i *TextInput) SetPreeditStylingHandler(f TextInputPreeditStylingHandlerFunc) {
	i.preeditStylingHandler = f
}

// TextInputPreeditCursorEvent : pre-edit cursor
//
// Sets the cursor position inside the composing text (as byte
// offset) relative to the start of the composing text. When index is a
// negative number no cursor is shown.
//
// This event is handled as part of a following preedit_string event.
type TextInputPreeditCursorEvent struct {
	Index int32
}
type TextInputPreeditCursorHandlerFunc func(TextInputPreeditCursorEvent)

// SetPreeditCursorHandler : sets handler for TextInputPreeditCursorEvent
func (i *TextInput) SetPreeditCursorHandler(f TextInputPreeditCursorHandlerFunc) {
	i.preeditCursorHandler = f
}

// TextInputCommitStringEvent : commit
//
// Notify when text should be inserted into the editor widget. The text to
// commit could be either just a single character after a key press or the
// result of some composing (pre-edit). It could also be an empty text
// when some text should be removed (see delete_surrounding_text) or when
// the input cursor should be moved (see cursor_position).
//
// Any previously set composing text should be removed.
type TextInputCommitStringEvent struct {
	Serial uint32
	Text   string
}
type TextInputCommitStringHandlerFunc func(TextInputCommitStringEvent)

// SetCommitStringHandler : sets handler for TextInputCommitStringEvent
func (i *TextInput) SetCommitStringHandler(f TextInputCommitStringHandlerFunc) {
	i.commitStringHandler = f
}

// TextInputCursorPositionEvent : set cursor to new position
//
// Notify when the cursor or anchor position should be modified.
//
// This event should be handled as part of a following commit_string
// event.
type TextInputCursorPositionEvent struct {
	Index  int32
	Anchor int32
}
type TextInputCursorPositionHandlerFunc func(TextInputCursorPositionEvent)

// SetCursorPositionHandler : sets handler for TextInputCursorPositionEvent
func (i *TextInput) SetCursorPositionHandler(f TextInputCursorPositionHandlerFunc) {
	i.cursorPositionHandler = f
}

// TextInputDeleteSurroundingTextEvent : delete surrounding text
//
// Notify when the text around the current cursor position should be
// deleted.
//
// Index is relative to the current cursor (in bytes).
// Length is the length of deleted text (in bytes).
//
// This event should be handled as part of a following commit_string
// event.
type TextInputDeleteSurroundingTextEvent struct {
	Index  int32
	Length uint32
}
type TextInputDeleteSurroundingTextHandlerFunc func(TextInputDeleteSurroundingTextEvent)

// SetDeleteSurroundingTextHandler : sets handler for TextInputDeleteSurroundingTextEvent
func (i *TextInput) SetDeleteSurroundingTextHandler(f TextInputDeleteSurroundingTextHandlerFunc) {
	i.deleteSurroundingTextHandler = f
}

// TextInputKeysymEvent : keysym
//
// Notify when a key event was sent. Key events should not be used
// for normal text input operations, which should be done with
// commit_string, delete_surrounding_text, etc. The key event follows
// the wl_keyboard key event convention. Sym is an XKB keysym, state a
// wl_keyboard key_state. Modifiers are a mask for effective modifiers
// (where the modifier indices are set by the modifiers_map event)
type TextInputKeysymEvent struct {
	Serial    uint32
	Time      uint32
	Sym       uint32
	State     uint32
	Modifiers uint32
}
type TextInputKeysymHandlerFunc func(TextInputKeysymEvent)

// SetKeysymHandler : sets handler for TextInputKeysymEvent
func (i *TextInput) SetKeysymHandler(f TextInputKeysymHandlerFunc) {
	i.keysymHandler = f
}

// TextInputLanguageEvent : language
//
// Sets the language of the input text. The "language" argument is an
// RFC-3066 format language tag.
type TextInputLanguageEvent struct {
	Serial   uint32
	Language string
}
type TextInputLanguageHandlerFunc func(TextInputLanguageEvent)

// SetLanguageHandler : sets handler for TextInputLanguageEvent
func (i *TextInput) SetLanguageHandler(f TextInputLanguageHandlerFunc) {
	i.languageHandler = f
}

// TextInputTextDirectionEvent : text direction
//
// Sets the text direction of input text.
//
// It is mainly needed for showing an input cursor on the correct side of
// the editor when there is no input done yet and making sure neutral
// direction text is laid out properly.
type TextInputTextDirectionEvent struct {
	Serial    uint32
	Direction uint32
}
type TextInputTextDirectionHandlerFunc func(TextInputTextDirectionEvent)

// SetTextDirectionHandler : sets handler for TextInputTextDirectionEvent
func (i *TextInput) SetTextDirectionHandler(f TextInputTextDirectionHandlerFunc) {
	i.textDirectionHandler = f
}

func (i *TextInput) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.enterHandler == nil {
			return
		}
		var e TextInputEnterEvent
		l := 0
		e.Surface = i.Context().GetOrRegister(client.Uint32(data[l:l+4]), (*client.Surface)(nil)).(*client.Surface)

		i.enterHandler(e)
	case 1:
		if i.leaveHandler == nil {
			return
		}
		var e TextInputLeaveEvent

		i.leaveHandler(e)
	case 2:
		if i.modifiersMapHandler == nil {
			return
		}
		var e TextInputModifiersMapEvent
		l := 0
		_mapLen := int(client.Uint32(data[l : l+4]))
		l += 4
		e.Map = make([]byte, _mapLen)
		copy(e.Map, data[l:l+_mapLen])

		i.modifiersMapHandler(e)
	case 3:
		if i.inputPanelStateHandler == nil {
			return
		}
		var e TextInputInputPanelStateEvent
		l := 0
		e.State = client.Uint32(data[l : l+4])

		i.inputPanelStateHandler(e)
	case 4:
		if i.preeditStringHandler == nil {
			return
		}
		var e TextInputPreeditStringEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		textLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Text = client.String(data[l : l+textLen])
		l += textLen
		commitLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Commit = client.String(data[l : l+commitLen])

		i.preeditStringHandler(e)
	case 5:
		if i.preeditStylingHandler == nil {
			return
		}
		var e TextInputPreeditStylingEvent
		l := 0
		e.Index = client.Uint32(data[l : l+4])
		l += 4
		e.Length = client.Uint32(data[l : l+4])
		l += 4
		e.Style = client.Uint32(data[l : l+4])

		i.preeditStylingHandler(e)
	case 6:
		if i.preeditCursorHandler == nil {
			return
		}
		var e TextInputPreeditCursorEvent
		l := 0
		e.Index = int32(client.Uint32(data[l : l+4]))

		i.preeditCursorHandler(e)
	case 7:
		if i.commitStringHandler == nil {
			return
		}
		var e TextInputCommitStringEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		textLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Text = client.String(data[l : l+textLen])

		i.commitStringHandler(e)
	case 8:
		if i.cursorPositionHandler == nil {
			return
		}
		var e TextInputCursorPositionEvent
		l := 0
		e.Index = int32(client.Uint32(data[l : l+4]))
		l += 4
		e.Anchor = int32(client.Uint32(data[l : l+4]))

		i.cursorPositionHandler(e)
	case 9:
		if i.deleteSurroundingTextHandler == nil {
			return
		}
		var e TextInputDeleteSurroundingTextEvent
		l := 0
		e.Index = int32(client.Uint32(data[l : l+4]))
		l += 4
		e.Length = client.Uint32(data[l : l+4])

		i.deleteSurroundingTextHandler(e)
	case 10:
		if i.keysymHandler == nil {
			return
		}
		var e TextInputKeysymEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		e.Time = client.Uint32(data[l : l+4])
		l += 4
		e.Sym = client.Uint32(data[l : l+4])
		l += 4
		e.State = client.Uint32(data[l : l+4])
		l += 4
		e.Modifiers = client.Uint32(data[l : l+4])

		i.keysymHandler(e)
	case 11:
		if i.languageHandler == nil {
			return
		}
		var e TextInputLanguageEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		languageLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Language = client.String(data[l : l+languageLen])

		i.languageHandler(e)
	case 12:
		if i.textDirectionHandler == nil {
			return
		}
		var e TextInputTextDirectionEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		e.Direction = client.Uint32(data[l : l+4])

		i.textDirectionHandler(e)
	}
}

// TextInputManager : text input manager
//
// A factory for text_input objects. This object is a global singleton.
type TextInputManager struct {
	client.BaseProxy
}

// NewTextInputManager : text input manager
//
// A factory for text_input objects. This object is a global singleton.
func NewTextInputManager(ctx *client.Context) *TextInputManager {
	zwpTextInputManagerV1 := &TextInputManager{}
	ctx.Register(zwpTextInputManagerV1)
	return zwpTextInputManagerV1
}

// CreateTextInput : create text input
//
// Creates a new text_input object.
func (i *TextInputManager) CreateTextInput() (*TextInput, error) {
	id := NewTextInput(i.Context())
	const opcode = 0
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

func (i *TextInputManager) Destroy() error {
	i.Context().Unregister(i)
	return nil
}
