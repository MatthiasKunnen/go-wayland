// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/d10d18f3d49374d2e3eb96d63511f32795aab5f7/unstable/text-input/text-input-unstable-v1.xml
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

import (
	"sync"

	"github.com/rajveermalviya/go-wayland/client"
)

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
	mu                            sync.RWMutex
	enterHandlers                 []TextInputEnterHandler
	leaveHandlers                 []TextInputLeaveHandler
	modifiersMapHandlers          []TextInputModifiersMapHandler
	inputPanelStateHandlers       []TextInputInputPanelStateHandler
	preeditStringHandlers         []TextInputPreeditStringHandler
	preeditStylingHandlers        []TextInputPreeditStylingHandler
	preeditCursorHandlers         []TextInputPreeditCursorHandler
	commitStringHandlers          []TextInputCommitStringHandler
	cursorPositionHandlers        []TextInputCursorPositionHandler
	deleteSurroundingTextHandlers []TextInputDeleteSurroundingTextHandler
	keysymHandlers                []TextInputKeysymHandler
	languageHandlers              []TextInputLanguageHandler
	textDirectionHandlers         []TextInputTextDirectionHandler
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
//
func (i *TextInput) Activate(seat *client.Seat, surface *client.Surface) error {
	err := i.Context().SendRequest(i, 0, seat, surface)
	return err
}

// Deactivate : request deactivation
//
// Requests the text_input object to be deactivated (typically when the
// text entry lost focus). The seat argument is a wl_seat which was used
// for activation.
//
func (i *TextInput) Deactivate(seat *client.Seat) error {
	err := i.Context().SendRequest(i, 1, seat)
	return err
}

// ShowInputPanel : show input panels
//
// Requests input panels (virtual keyboard) to show.
//
func (i *TextInput) ShowInputPanel() error {
	err := i.Context().SendRequest(i, 2)
	return err
}

// HideInputPanel : hide input panels
//
// Requests input panels (virtual keyboard) to hide.
//
func (i *TextInput) HideInputPanel() error {
	err := i.Context().SendRequest(i, 3)
	return err
}

// Reset : reset
//
// Should be called by an editor widget when the input state should be
// reset, for example after the text was changed outside of the normal
// input method flow.
//
func (i *TextInput) Reset() error {
	err := i.Context().SendRequest(i, 4)
	return err
}

// SetSurroundingText : sets the surrounding text
//
// Sets the plain surrounding text around the input position. Text is
// UTF-8 encoded. Cursor is the byte offset within the
// surrounding text. Anchor is the byte offset of the
// selection anchor within the surrounding text. If there is no selected
// text anchor, then it is the same as cursor.
//
func (i *TextInput) SetSurroundingText(text string, cursor, anchor uint32) error {
	err := i.Context().SendRequest(i, 5, text, cursor, anchor)
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
//
func (i *TextInput) SetContentType(hint, purpose uint32) error {
	err := i.Context().SendRequest(i, 6, hint, purpose)
	return err
}

// SetCursorRectangle :
//
func (i *TextInput) SetCursorRectangle(x, y, width, height int32) error {
	err := i.Context().SendRequest(i, 7, x, y, width, height)
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
//
func (i *TextInput) SetPreferredLanguage(language string) error {
	err := i.Context().SendRequest(i, 8, language)
	return err
}

// CommitState :
//
//  serial: used to identify the known state
func (i *TextInput) CommitState(serial uint32) error {
	err := i.Context().SendRequest(i, 9, serial)
	return err
}

// InvokeAction :
//
func (i *TextInput) InvokeAction(button, index uint32) error {
	err := i.Context().SendRequest(i, 10, button, index)
	return err
}

func (i *TextInput) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

// TextInputContentHint : content hint
//
// Content hint is a bitmask to allow to modify the behavior of the text
// input.
const (
	// TextInputContentHintNone : no special behaviour
	TextInputContentHintNone = 0x0
	// TextInputContentHintDefault : auto completion, correction and capitalization
	TextInputContentHintDefault = 0x7
	// TextInputContentHintPassword : hidden and sensitive text
	TextInputContentHintPassword = 0xc0
	// TextInputContentHintAutoCompletion : suggest word completions
	TextInputContentHintAutoCompletion = 0x1
	// TextInputContentHintAutoCorrection : suggest word corrections
	TextInputContentHintAutoCorrection = 0x2
	// TextInputContentHintAutoCapitalization : switch to uppercase letters at the start of a sentence
	TextInputContentHintAutoCapitalization = 0x4
	// TextInputContentHintLowercase : prefer lowercase letters
	TextInputContentHintLowercase = 0x8
	// TextInputContentHintUppercase : prefer uppercase letters
	TextInputContentHintUppercase = 0x10
	// TextInputContentHintTitlecase : prefer casing for titles and headings (can be language dependent)
	TextInputContentHintTitlecase = 0x20
	// TextInputContentHintHiddenText : characters should be hidden
	TextInputContentHintHiddenText = 0x40
	// TextInputContentHintSensitiveData : typed text should not be stored
	TextInputContentHintSensitiveData = 0x80
	// TextInputContentHintLatin : just latin characters should be entered
	TextInputContentHintLatin = 0x100
	// TextInputContentHintMultiline : the text input is multiline
	TextInputContentHintMultiline = 0x200
)

// TextInputContentPurpose : content purpose
//
// The content purpose allows to specify the primary purpose of a text
// input.
//
// This allows an input method to show special purpose input panels with
// extra characters or to disallow some characters.
const (
	// TextInputContentPurposeNormal : default input, allowing all characters
	TextInputContentPurposeNormal = 0
	// TextInputContentPurposeAlpha : allow only alphabetic characters
	TextInputContentPurposeAlpha = 1
	// TextInputContentPurposeDigits : allow only digits
	TextInputContentPurposeDigits = 2
	// TextInputContentPurposeNumber : input a number (including decimal separator and sign)
	TextInputContentPurposeNumber = 3
	// TextInputContentPurposePhone : input a phone number
	TextInputContentPurposePhone = 4
	// TextInputContentPurposeUrl : input an URL
	TextInputContentPurposeUrl = 5
	// TextInputContentPurposeEmail : input an email address
	TextInputContentPurposeEmail = 6
	// TextInputContentPurposeName : input a name of a person
	TextInputContentPurposeName = 7
	// TextInputContentPurposePassword : input a password (combine with password or sensitive_data hint)
	TextInputContentPurposePassword = 8
	// TextInputContentPurposeDate : input a date
	TextInputContentPurposeDate = 9
	// TextInputContentPurposeTime : input a time
	TextInputContentPurposeTime = 10
	// TextInputContentPurposeDatetime : input a date and time
	TextInputContentPurposeDatetime = 11
	// TextInputContentPurposeTerminal : input for a terminal
	TextInputContentPurposeTerminal = 12
)

// TextInputPreeditStyle :
const (
	// TextInputPreeditStyleDefault : default style for composing text
	TextInputPreeditStyleDefault = 0
	// TextInputPreeditStyleNone : style should be the same as in non-composing text
	TextInputPreeditStyleNone      = 1
	TextInputPreeditStyleActive    = 2
	TextInputPreeditStyleInactive  = 3
	TextInputPreeditStyleHighlight = 4
	TextInputPreeditStyleUnderline = 5
	TextInputPreeditStyleSelection = 6
	TextInputPreeditStyleIncorrect = 7
)

// TextInputTextDirection :
const (
	// TextInputTextDirectionAuto : automatic text direction based on text and language
	TextInputTextDirectionAuto = 0
	// TextInputTextDirectionLtr : left-to-right
	TextInputTextDirectionLtr = 1
	// TextInputTextDirectionRtl : right-to-left
	TextInputTextDirectionRtl = 2
)

// TextInputEnterEvent : enter event
//
// Notify the text_input object when it received focus. Typically in
// response to an activate request.
type TextInputEnterEvent struct {
	Surface *client.Surface
}

type TextInputEnterHandler interface {
	HandleTextInputEnter(TextInputEnterEvent)
}

// AddEnterHandler : adds handler for TextInputEnterEvent
func (i *TextInput) AddEnterHandler(h TextInputEnterHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.enterHandlers = append(i.enterHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveEnterHandler(h TextInputEnterHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.enterHandlers {
		if e == h {
			i.enterHandlers = append(i.enterHandlers[:j], i.enterHandlers[j+1:]...)
			break
		}
	}
}

// TextInputLeaveEvent : leave event
//
// Notify the text_input object when it lost focus. Either in response
// to a deactivate request or when the assigned surface lost focus or was
// destroyed.
type TextInputLeaveEvent struct{}
type TextInputLeaveHandler interface {
	HandleTextInputLeave(TextInputLeaveEvent)
}

// AddLeaveHandler : adds handler for TextInputLeaveEvent
func (i *TextInput) AddLeaveHandler(h TextInputLeaveHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.leaveHandlers = append(i.leaveHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveLeaveHandler(h TextInputLeaveHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.leaveHandlers {
		if e == h {
			i.leaveHandlers = append(i.leaveHandlers[:j], i.leaveHandlers[j+1:]...)
			break
		}
	}
}

// TextInputModifiersMapEvent : modifiers map
//
// Transfer an array of 0-terminated modifier names. The position in
// the array is the index of the modifier as used in the modifiers
// bitmask in the keysym event.
type TextInputModifiersMapEvent struct {
	Map []int32
}

type TextInputModifiersMapHandler interface {
	HandleTextInputModifiersMap(TextInputModifiersMapEvent)
}

// AddModifiersMapHandler : adds handler for TextInputModifiersMapEvent
func (i *TextInput) AddModifiersMapHandler(h TextInputModifiersMapHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.modifiersMapHandlers = append(i.modifiersMapHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveModifiersMapHandler(h TextInputModifiersMapHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.modifiersMapHandlers {
		if e == h {
			i.modifiersMapHandlers = append(i.modifiersMapHandlers[:j], i.modifiersMapHandlers[j+1:]...)
			break
		}
	}
}

// TextInputInputPanelStateEvent : state of the input panel
//
// Notify when the visibility state of the input panel changed.
type TextInputInputPanelStateEvent struct {
	State uint32
}

type TextInputInputPanelStateHandler interface {
	HandleTextInputInputPanelState(TextInputInputPanelStateEvent)
}

// AddInputPanelStateHandler : adds handler for TextInputInputPanelStateEvent
func (i *TextInput) AddInputPanelStateHandler(h TextInputInputPanelStateHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.inputPanelStateHandlers = append(i.inputPanelStateHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveInputPanelStateHandler(h TextInputInputPanelStateHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.inputPanelStateHandlers {
		if e == h {
			i.inputPanelStateHandlers = append(i.inputPanelStateHandlers[:j], i.inputPanelStateHandlers[j+1:]...)
			break
		}
	}
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

type TextInputPreeditStringHandler interface {
	HandleTextInputPreeditString(TextInputPreeditStringEvent)
}

// AddPreeditStringHandler : adds handler for TextInputPreeditStringEvent
func (i *TextInput) AddPreeditStringHandler(h TextInputPreeditStringHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.preeditStringHandlers = append(i.preeditStringHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemovePreeditStringHandler(h TextInputPreeditStringHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.preeditStringHandlers {
		if e == h {
			i.preeditStringHandlers = append(i.preeditStringHandlers[:j], i.preeditStringHandlers[j+1:]...)
			break
		}
	}
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

type TextInputPreeditStylingHandler interface {
	HandleTextInputPreeditStyling(TextInputPreeditStylingEvent)
}

// AddPreeditStylingHandler : adds handler for TextInputPreeditStylingEvent
func (i *TextInput) AddPreeditStylingHandler(h TextInputPreeditStylingHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.preeditStylingHandlers = append(i.preeditStylingHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemovePreeditStylingHandler(h TextInputPreeditStylingHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.preeditStylingHandlers {
		if e == h {
			i.preeditStylingHandlers = append(i.preeditStylingHandlers[:j], i.preeditStylingHandlers[j+1:]...)
			break
		}
	}
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

type TextInputPreeditCursorHandler interface {
	HandleTextInputPreeditCursor(TextInputPreeditCursorEvent)
}

// AddPreeditCursorHandler : adds handler for TextInputPreeditCursorEvent
func (i *TextInput) AddPreeditCursorHandler(h TextInputPreeditCursorHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.preeditCursorHandlers = append(i.preeditCursorHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemovePreeditCursorHandler(h TextInputPreeditCursorHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.preeditCursorHandlers {
		if e == h {
			i.preeditCursorHandlers = append(i.preeditCursorHandlers[:j], i.preeditCursorHandlers[j+1:]...)
			break
		}
	}
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

type TextInputCommitStringHandler interface {
	HandleTextInputCommitString(TextInputCommitStringEvent)
}

// AddCommitStringHandler : adds handler for TextInputCommitStringEvent
func (i *TextInput) AddCommitStringHandler(h TextInputCommitStringHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.commitStringHandlers = append(i.commitStringHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveCommitStringHandler(h TextInputCommitStringHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.commitStringHandlers {
		if e == h {
			i.commitStringHandlers = append(i.commitStringHandlers[:j], i.commitStringHandlers[j+1:]...)
			break
		}
	}
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

type TextInputCursorPositionHandler interface {
	HandleTextInputCursorPosition(TextInputCursorPositionEvent)
}

// AddCursorPositionHandler : adds handler for TextInputCursorPositionEvent
func (i *TextInput) AddCursorPositionHandler(h TextInputCursorPositionHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.cursorPositionHandlers = append(i.cursorPositionHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveCursorPositionHandler(h TextInputCursorPositionHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.cursorPositionHandlers {
		if e == h {
			i.cursorPositionHandlers = append(i.cursorPositionHandlers[:j], i.cursorPositionHandlers[j+1:]...)
			break
		}
	}
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

type TextInputDeleteSurroundingTextHandler interface {
	HandleTextInputDeleteSurroundingText(TextInputDeleteSurroundingTextEvent)
}

// AddDeleteSurroundingTextHandler : adds handler for TextInputDeleteSurroundingTextEvent
func (i *TextInput) AddDeleteSurroundingTextHandler(h TextInputDeleteSurroundingTextHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.deleteSurroundingTextHandlers = append(i.deleteSurroundingTextHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveDeleteSurroundingTextHandler(h TextInputDeleteSurroundingTextHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.deleteSurroundingTextHandlers {
		if e == h {
			i.deleteSurroundingTextHandlers = append(i.deleteSurroundingTextHandlers[:j], i.deleteSurroundingTextHandlers[j+1:]...)
			break
		}
	}
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

type TextInputKeysymHandler interface {
	HandleTextInputKeysym(TextInputKeysymEvent)
}

// AddKeysymHandler : adds handler for TextInputKeysymEvent
func (i *TextInput) AddKeysymHandler(h TextInputKeysymHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.keysymHandlers = append(i.keysymHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveKeysymHandler(h TextInputKeysymHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.keysymHandlers {
		if e == h {
			i.keysymHandlers = append(i.keysymHandlers[:j], i.keysymHandlers[j+1:]...)
			break
		}
	}
}

// TextInputLanguageEvent : language
//
// Sets the language of the input text. The "language" argument is an
// RFC-3066 format language tag.
type TextInputLanguageEvent struct {
	Serial   uint32
	Language string
}

type TextInputLanguageHandler interface {
	HandleTextInputLanguage(TextInputLanguageEvent)
}

// AddLanguageHandler : adds handler for TextInputLanguageEvent
func (i *TextInput) AddLanguageHandler(h TextInputLanguageHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.languageHandlers = append(i.languageHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveLanguageHandler(h TextInputLanguageHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.languageHandlers {
		if e == h {
			i.languageHandlers = append(i.languageHandlers[:j], i.languageHandlers[j+1:]...)
			break
		}
	}
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

type TextInputTextDirectionHandler interface {
	HandleTextInputTextDirection(TextInputTextDirectionEvent)
}

// AddTextDirectionHandler : adds handler for TextInputTextDirectionEvent
func (i *TextInput) AddTextDirectionHandler(h TextInputTextDirectionHandler) {
	if h == nil {
		return
	}

	i.mu.Lock()
	i.textDirectionHandlers = append(i.textDirectionHandlers, h)
	i.mu.Unlock()
}

func (i *TextInput) RemoveTextDirectionHandler(h TextInputTextDirectionHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for j, e := range i.textDirectionHandlers {
		if e == h {
			i.textDirectionHandlers = append(i.textDirectionHandlers[:j], i.textDirectionHandlers[j+1:]...)
			break
		}
	}
}

func (i *TextInput) Dispatch(event *client.Event) {
	switch event.Opcode {
	case 0:
		i.mu.RLock()
		if len(i.enterHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputEnterEvent{
			Surface: event.Proxy(i.Context()).(*client.Surface),
		}

		i.mu.RLock()
		for _, h := range i.enterHandlers {
			i.mu.RUnlock()

			h.HandleTextInputEnter(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 1:
		i.mu.RLock()
		if len(i.leaveHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputLeaveEvent{}

		i.mu.RLock()
		for _, h := range i.leaveHandlers {
			i.mu.RUnlock()

			h.HandleTextInputLeave(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 2:
		i.mu.RLock()
		if len(i.modifiersMapHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputModifiersMapEvent{
			Map: event.Array(),
		}

		i.mu.RLock()
		for _, h := range i.modifiersMapHandlers {
			i.mu.RUnlock()

			h.HandleTextInputModifiersMap(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 3:
		i.mu.RLock()
		if len(i.inputPanelStateHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputInputPanelStateEvent{
			State: event.Uint32(),
		}

		i.mu.RLock()
		for _, h := range i.inputPanelStateHandlers {
			i.mu.RUnlock()

			h.HandleTextInputInputPanelState(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 4:
		i.mu.RLock()
		if len(i.preeditStringHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputPreeditStringEvent{
			Serial: event.Uint32(),
			Text:   event.String(),
			Commit: event.String(),
		}

		i.mu.RLock()
		for _, h := range i.preeditStringHandlers {
			i.mu.RUnlock()

			h.HandleTextInputPreeditString(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 5:
		i.mu.RLock()
		if len(i.preeditStylingHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputPreeditStylingEvent{
			Index:  event.Uint32(),
			Length: event.Uint32(),
			Style:  event.Uint32(),
		}

		i.mu.RLock()
		for _, h := range i.preeditStylingHandlers {
			i.mu.RUnlock()

			h.HandleTextInputPreeditStyling(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 6:
		i.mu.RLock()
		if len(i.preeditCursorHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputPreeditCursorEvent{
			Index: event.Int32(),
		}

		i.mu.RLock()
		for _, h := range i.preeditCursorHandlers {
			i.mu.RUnlock()

			h.HandleTextInputPreeditCursor(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 7:
		i.mu.RLock()
		if len(i.commitStringHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputCommitStringEvent{
			Serial: event.Uint32(),
			Text:   event.String(),
		}

		i.mu.RLock()
		for _, h := range i.commitStringHandlers {
			i.mu.RUnlock()

			h.HandleTextInputCommitString(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 8:
		i.mu.RLock()
		if len(i.cursorPositionHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputCursorPositionEvent{
			Index:  event.Int32(),
			Anchor: event.Int32(),
		}

		i.mu.RLock()
		for _, h := range i.cursorPositionHandlers {
			i.mu.RUnlock()

			h.HandleTextInputCursorPosition(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 9:
		i.mu.RLock()
		if len(i.deleteSurroundingTextHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputDeleteSurroundingTextEvent{
			Index:  event.Int32(),
			Length: event.Uint32(),
		}

		i.mu.RLock()
		for _, h := range i.deleteSurroundingTextHandlers {
			i.mu.RUnlock()

			h.HandleTextInputDeleteSurroundingText(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 10:
		i.mu.RLock()
		if len(i.keysymHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputKeysymEvent{
			Serial:    event.Uint32(),
			Time:      event.Uint32(),
			Sym:       event.Uint32(),
			State:     event.Uint32(),
			Modifiers: event.Uint32(),
		}

		i.mu.RLock()
		for _, h := range i.keysymHandlers {
			i.mu.RUnlock()

			h.HandleTextInputKeysym(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 11:
		i.mu.RLock()
		if len(i.languageHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputLanguageEvent{
			Serial:   event.Uint32(),
			Language: event.String(),
		}

		i.mu.RLock()
		for _, h := range i.languageHandlers {
			i.mu.RUnlock()

			h.HandleTextInputLanguage(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
	case 12:
		i.mu.RLock()
		if len(i.textDirectionHandlers) == 0 {
			i.mu.RUnlock()
			break
		}
		i.mu.RUnlock()

		e := TextInputTextDirectionEvent{
			Serial:    event.Uint32(),
			Direction: event.Uint32(),
		}

		i.mu.RLock()
		for _, h := range i.textDirectionHandlers {
			i.mu.RUnlock()

			h.HandleTextInputTextDirection(e)

			i.mu.RLock()
		}
		i.mu.RUnlock()
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
//
func (i *TextInputManager) CreateTextInput() (*TextInput, error) {
	id := NewTextInput(i.Context())
	err := i.Context().SendRequest(i, 0, id)
	return id, err
}

func (i *TextInputManager) Destroy() error {
	i.Context().Unregister(i)
	return nil
}
