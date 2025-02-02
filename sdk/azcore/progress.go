//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

import (
	"io"
)

type progress struct {
	rc     io.ReadCloser
	rsc    ReadSeekCloser
	pr     func(bytesTransferred int64)
	offset int64
}

// NewRequestProgress adds progress reporting to an HTTP request's body stream.
func NewRequestProgress(body ReadSeekCloser, pr func(bytesTransferred int64)) ReadSeekCloser {
	return &progress{
		rc:     body,
		rsc:    body,
		pr:     pr,
		offset: 0,
	}
}

// NewResponseProgress adds progress reporting to an HTTP response's body stream.
func NewResponseProgress(body io.ReadCloser, pr func(bytesTransferred int64)) io.ReadCloser {
	return &progress{
		rc:     body,
		rsc:    nil,
		pr:     pr,
		offset: 0,
	}
}

// Read reads a block of data from an inner stream and reports progress
func (p *progress) Read(b []byte) (n int, err error) {
	n, err = p.rc.Read(b)
	if err != nil && err != io.EOF {
		return
	}
	p.offset += int64(n)
	// Invokes the user's callback method to report progress
	p.pr(p.offset)
	return
}

// Seek only expects a zero or from beginning.
func (p *progress) Seek(offset int64, whence int) (int64, error) {
	// This should only ever be called with offset = 0 and whence = io.SeekStart
	n, err := p.rsc.Seek(offset, whence)
	if err == nil {
		p.offset = int64(n)
	}
	return n, err
}

// requestBodyProgress supports Close but the underlying stream may not; if it does, Close will close it.
func (p *progress) Close() error {
	return p.rc.Close()
}
