package whisper

import (
	"fmt"
)

///////////////////////////////////////////////////////////////////////////////
// CGO

/*
#include <whisper.h>
*/
import "C"

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *Params_C) SetTranslate(v bool) {
	p.translate = toBool(v)
}

func (p *Params_C) SetSplitOnWord(v bool) {
	p.split_on_word = toBool(v)
}

func (p *Params_C) SetNoContext(v bool) {
	p.no_context = toBool(v)
}

func (p *Params_C) SetSingleSegment(v bool) {
	p.single_segment = toBool(v)
}

func (p *Params_C) SetPrintSpecial(v bool) {
	p.print_special = toBool(v)
}

func (p *Params_C) SetPrintProgress(v bool) {
	p.print_progress = toBool(v)
}

func (p *Params_C) SetPrintRealtime(v bool) {
	p.print_realtime = toBool(v)
}

func (p *Params_C) SetPrintTimestamps(v bool) {
	p.print_timestamps = toBool(v)
}

// Set language id
func (p *Params_C) SetLanguage(lang int) error {
	if lang == -1 {
		p.language = nil
		return nil
	}
	str := C.whisper_lang_str(C.int(lang))
	if str == nil {
		return ErrInvalidLanguage
	} else {
		p.language = str
	}
	return nil
}

// Get language id
func (p *Params_C) Language() int {
	if p.language == nil {
		return -1
	}
	return int(C.whisper_lang_id(p.language))
}

// Threads available
func (p *Params_C) Threads() int {
	return int(p.n_threads)
}

// Set number of threads to use
func (p *Params_C) SetThreads(threads int) {
	p.n_threads = C.int(threads)
}

// Set start offset in ms
func (p *Params_C) SetOffset(offset_ms int) {
	p.offset_ms = C.int(offset_ms)
}

// Set audio duration to process in ms
func (p *Params_C) SetDuration(duration_ms int) {
	p.duration_ms = C.int(duration_ms)
}

// Set timestamp token probability threshold (~0.01)
func (p *Params_C) SetTokenThreshold(t float32) {
	p.thold_pt = C.float(t)
}

// Set timestamp token sum probability threshold (~0.01)
func (p *Params_C) SetTokenSumThreshold(t float32) {
	p.thold_ptsum = C.float(t)
}

// Set max segment length in characters
func (p *Params_C) SetMaxSegmentLength(n int) {
	p.max_len = C.int(n)
}

func (p *Params_C) SetTokenTimestamps(b bool) {
	p.token_timestamps = toBool(b)
}

// Set max tokens per segment (0 = no limit)
func (p *Params_C) SetMaxTokensPerSegment(n int) {
	p.max_tokens = C.int(n)
}

// Set audio encoder context
func (p *Params_C) SetAudioCtx(n int) {
	p.audio_ctx = C.int(n)
}

// Set initial prompt
func (p *Params_C) SetInitialPrompt(prompt string) {
	p.initial_prompt = C.CString(prompt)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func toBool(v bool) C.bool {
	if v {
		return C.bool(true)
	}
	return C.bool(false)
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (p *Params_C) String() string {
	str := "<whisper.params"
	str += fmt.Sprintf(" strategy=%v", p.strategy)
	str += fmt.Sprintf(" n_threads=%d", p.n_threads)
	if p.language != nil {
		str += fmt.Sprintf(" language=%s", C.GoString(p.language))
	}
	str += fmt.Sprintf(" n_max_text_ctx=%d", p.n_max_text_ctx)
	str += fmt.Sprintf(" offset_ms=%d", p.offset_ms)
	str += fmt.Sprintf(" duration_ms=%d", p.duration_ms)
	str += fmt.Sprintf(" audio_ctx=%d", p.audio_ctx)
	str += fmt.Sprintf(" initial_prompt=%s", C.GoString(p.initial_prompt))
	if p.translate {
		str += " translate"
	}
	if p.no_context {
		str += " no_context"
	}
	if p.single_segment {
		str += " single_segment"
	}
	if p.print_special {
		str += " print_special"
	}
	if p.print_progress {
		str += " print_progress"
	}
	if p.print_realtime {
		str += " print_realtime"
	}
	if p.print_timestamps {
		str += " print_timestamps"
	}
	if p.token_timestamps {
		str += " token_timestamps"
	}

	return str + ">"
}
