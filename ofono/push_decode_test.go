/*
 * Copyright 2014 Canonical Ltd.
 *
 * Authors:
 * Sergio Schvezov: sergio.schvezov@canonical.com
 *
 * This file is part of mms.
 *
 * mms is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * mms is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package ofono

import (
	"errors"
	"testing"

	"github.com/ubuntu-phonedations/nuntium/mms"
	. "launchpad.net/gocheck"
)

type PushDecodeTestSuite struct {
	pdu *PushPDU
}

var _ = Suite(&PushDecodeTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *PushDecodeTestSuite) SetUpTest(c *C) {
	s.pdu = new(PushPDU)
}

func (s *PushDecodeTestSuite) TestDecodeVodaphoneSpain(c *C) {
	inputBytes := []byte{
		0x00, 0x06, 0x26, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
		0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
		0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0xb4, 0x81,
		0x8d, 0xdf, 0x8c, 0x82, 0x98, 0x4e, 0x4f, 0x4b, 0x35, 0x43, 0x64, 0x7a, 0x30,
		0x38, 0x42, 0x41, 0x73, 0x77, 0x61, 0x62, 0x77, 0x55, 0x48, 0x00, 0x8d, 0x90,
		0x89, 0x18, 0x80, 0x2b, 0x33, 0x34, 0x36, 0x30, 0x30, 0x39, 0x34, 0x34, 0x34,
		0x36, 0x33, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00,
		0x8a, 0x80, 0x8e, 0x02, 0x74, 0x00, 0x88, 0x05, 0x81, 0x03, 0x02, 0xa3, 0x00,
		0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6d, 0x6d, 0x31, 0x66, 0x65,
		0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x6c, 0x65, 0x74, 0x73, 0x2f, 0x4e, 0x4f,
		0x4b, 0x35, 0x43, 0x64, 0x7a, 0x30, 0x38, 0x42, 0x41, 0x73, 0x77, 0x61, 0x62,
		0x77, 0x55, 0x48, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 38)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 106)
}

func (s *PushDecodeTestSuite) TestDecodeTelecomPersonal(c *C) {
	inputBytes := []byte{
		0x01, 0x06, 0x26, 0x61, 0x70, 0x70, 0x6C, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
		0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
		0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0xb4, 0x86,
		0xc3, 0x95, 0x8c, 0x82, 0x98, 0x6d, 0x30, 0x34, 0x42, 0x4b, 0x6b, 0x73, 0x69,
		0x6d, 0x30, 0x35, 0x40, 0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
		0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x00, 0x8d, 0x90,
		0x89, 0x19, 0x80, 0x2b, 0x35, 0x34, 0x33, 0x35, 0x31, 0x35, 0x39, 0x32, 0x34,
		0x39, 0x30, 0x36, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e,
		0x00, 0x8a, 0x80, 0x8e, 0x02, 0x74, 0x00, 0x88, 0x05, 0x81, 0x03, 0x02, 0xa2,
		0xff, 0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x37, 0x32, 0x2e,
		0x32, 0x35, 0x2e, 0x37, 0x2e, 0x31, 0x33, 0x31, 0x2f, 0x3f, 0x6d, 0x65, 0x73,
		0x73, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x64, 0x3d, 0x6d, 0x30, 0x34, 0x42, 0x4b,
		0x68, 0x34, 0x33, 0x65, 0x30, 0x33, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 38)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 122)
}

func (s *PushDecodeTestSuite) TestDecodeATTUSA(c *C) {
	inputBytes := []byte{
		0x01, 0x06, 0x27, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
		0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d, 0x6d, 0x73,
		0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf, 0x84, 0x8d, 0x01,
		0x82, 0xb4, 0x84, 0x8c, 0x82, 0x98, 0x44, 0x32, 0x30, 0x34, 0x30, 0x37, 0x31,
		0x36, 0x35, 0x36, 0x32, 0x34, 0x36, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x00, 0x8d, 0x90, 0x89, 0x18, 0x80,
		0x2b, 0x31, 0x37, 0x37, 0x34, 0x32, 0x37, 0x30, 0x30, 0x36, 0x35, 0x39, 0x2f,
		0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00, 0x96, 0x02, 0xea,
		0x00, 0x8a, 0x80, 0x8e, 0x02, 0x80, 0x00, 0x88, 0x05, 0x81, 0x03, 0x05, 0x46,
		0x00, 0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x36, 0x36, 0x2e,
		0x32, 0x31, 0x36, 0x2e, 0x31, 0x36, 0x36, 0x2e, 0x36, 0x37, 0x3a, 0x38, 0x30,
		0x30, 0x34, 0x2f, 0x30, 0x34, 0x30, 0x37, 0x31, 0x36, 0x35, 0x36, 0x32, 0x34,
		0x36, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 39)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 130)
}

func (s *PushDecodeTestSuite) TestDecodeSoneraFinland(c *C) {
	inputBytes := []byte{
		0x00, 0x06, 0x07, 0xbe, 0xaf, 0x84, 0x8d, 0xf2, 0xb4, 0x81, 0x8c, 0x82, 0x98,
		0x41, 0x42, 0x73, 0x54, 0x4c, 0x4e, 0x41, 0x4c, 0x41, 0x6d, 0x6d, 0x4e, 0x33,
		0x77, 0x72, 0x38, 0x32, 0x00, 0x8d, 0x92, 0x89, 0x19, 0x80, 0x2b, 0x33, 0x35,
		0x38, 0x34, 0x30, 0x37, 0x36, 0x39, 0x34, 0x34, 0x38, 0x34, 0x2f, 0x54, 0x59,
		0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00, 0x86, 0x81, 0x8a, 0x80, 0x8e,
		0x03, 0x03, 0x15, 0x85, 0x88, 0x05, 0x81, 0x03, 0x03, 0xf4, 0x7f, 0x83, 0x68,
		0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6d, 0x6d, 0x73, 0x63, 0x36, 0x31, 0x3a,
		0x31, 0x30, 0x30, 0x32, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x63, 0x2f, 0x36, 0x5f,
		0x31, 0x3f, 0x41, 0x42, 0x73, 0x54, 0x4c, 0x4e, 0x41, 0x4c, 0x41, 0x6d, 0x6d,
		0x4e, 0x33, 0x77, 0x72, 0x38, 0x32, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 7)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 114)
}

func (s *PushDecodeTestSuite) TestOperatorWithContentLength(c *C) {
	inputBytes := []byte{
		0x01, 0x06, 0x07, 0xbe, 0x8d, 0xf0, 0xaf, 0x84, 0xb4, 0x84, 0x8c, 0x82, 0x98,
		0x41, 0x78, 0x67, 0x41, 0x6a, 0x45, 0x73, 0x49, 0x47, 0x46, 0x57, 0x45, 0x54,
		0x45, 0x53, 0x76, 0x41, 0x00, 0x8d, 0x93, 0x89, 0x18, 0x80, 0x2b, 0x33, 0x31,
		0x36, 0x35, 0x35, 0x35, 0x38, 0x34, 0x34, 0x32, 0x35, 0x2f, 0x54, 0x59, 0x50,
		0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00, 0x86, 0x81, 0x8a, 0x80, 0x8e, 0x03,
		0x01, 0xc5, 0x0d, 0x88, 0x05, 0x81, 0x03, 0x03, 0xf4, 0x80, 0x83, 0x68, 0x74,
		0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6d, 0x70, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x65,
		0x6c, 0x2e, 0x6b, 0x70, 0x6e, 0x2f, 0x6d, 0x6d, 0x73, 0x63, 0x2f, 0x30, 0x31,
		0x3f, 0x41, 0x78, 0x67, 0x41, 0x6a, 0x45, 0x73, 0x49, 0x47, 0x46, 0x57, 0x45,
		0x54, 0x45, 0x53, 0x76, 0x41, 0x00,
	}

	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 7)
	c.Check(int(s.pdu.ContentLength), Equals, 112)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 113)
}

func (s *PushDecodeTestSuite) TestDecodeNonPushPDU(c *C) {
	inputBytes := []byte{
		0x00, 0x07, 0x07, 0xbe, 0xaf, 0x84, 0x8d, 0xf2, 0xb4, 0x81, 0x8c,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), DeepEquals, errors.New("7 != 6 is not a push PDU"))
}

func (s *PushDecodeTestSuite) TestDecodeTMobileUSA(c *C) {
	inputBytes := []byte{
		0xc0, 0x06, 0x28, 0x1f, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
		0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d,
		0x6d, 0x73, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0x81, 0x84,
		0x8d, 0x80, 0xaf, 0x84, 0x8c, 0x82, 0x98, 0x6d, 0x61, 0x76, 0x6f, 0x64, 0x69,
		0x2d, 0x37, 0x2d, 0x38, 0x39, 0x2d, 0x31, 0x63, 0x30, 0x2d, 0x37, 0x2d, 0x63,
		0x61, 0x2d, 0x35, 0x30, 0x66, 0x39, 0x33, 0x38, 0x34, 0x33, 0x2d, 0x37, 0x2d,
		0x31, 0x33, 0x62, 0x2d, 0x32, 0x65, 0x62, 0x2d, 0x31, 0x2d, 0x63, 0x61, 0x2d,
		0x33, 0x36, 0x31, 0x65, 0x33, 0x31, 0x35, 0x00, 0x8d, 0x92, 0x89, 0x1a, 0x80,
		0x18, 0x83, 0x2b, 0x31, 0x39, 0x31, 0x39, 0x39, 0x30, 0x33, 0x33, 0x34, 0x38,
		0x38, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00, 0x8a,
		0x80, 0x8e, 0x03, 0x0f, 0x21, 0x9f, 0x88, 0x05, 0x81, 0x03, 0x03, 0xf4, 0x80,
		0x83, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x61, 0x74, 0x6c, 0x32, 0x6d,
		0x6f, 0x73, 0x67, 0x65, 0x74, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x65, 0x6e, 0x67,
		0x2e, 0x74, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
		0x2f, 0x6d, 0x6d, 0x73, 0x2f, 0x77, 0x61, 0x70, 0x65, 0x6e, 0x63, 0x3f, 0x54,
		0x3d, 0x6d, 0x61, 0x76, 0x6f, 0x64, 0x69, 0x2d, 0x37, 0x2d, 0x31, 0x33, 0x62,
		0x2d, 0x32, 0x65, 0x62, 0x2d, 0x31, 0x2d, 0x63, 0x61, 0x2d, 0x33, 0x36, 0x31,
		0x65, 0x33, 0x31, 0x35, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 40)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 183)
}

func (s *PushDecodeTestSuite) TestDecodePlayPoland(c *C) {
	inputBytes := []byte{
		0x2e, 0x06, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
		0x6f, 0x6e, 0x2f, 0x76, 0x6e, 0x64, 0x2e, 0x77, 0x61, 0x70, 0x2e, 0x6d,
		0x6d, 0x73, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x00, 0xaf,
		0x84, 0x8c, 0x82, 0x98, 0x31, 0x34, 0x34, 0x32, 0x34, 0x30, 0x31, 0x33,
		0x31, 0x38, 0x40, 0x6d, 0x6d, 0x73, 0x32, 0x00, 0x8d, 0x92, 0x89, 0x18,
		0x80, 0x2b, 0x34, 0x38, 0x38, 0x38, 0x32, 0x30, 0x34, 0x30, 0x32, 0x32,
		0x35, 0x2f, 0x54, 0x59, 0x50, 0x45, 0x3d, 0x50, 0x4c, 0x4d, 0x4e, 0x00,
		0x8f, 0x81, 0x86, 0x80, 0x8a, 0x80, 0x8e, 0x03, 0x03, 0xad, 0x21, 0x88,
		0x05, 0x81, 0x03, 0x03, 0xf4, 0x80, 0x83, 0x68, 0x74, 0x74, 0x70, 0x3a,
		0x2f, 0x2f, 0x6d, 0x6d, 0x73, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x2e,
		0x70, 0x6c, 0x2f, 0x3f, 0x69, 0x64, 0x3d, 0x31, 0x34, 0x34, 0x32, 0x34,
		0x30, 0x31, 0x33, 0x31, 0x38, 0x42, 0x00,
	}
	dec := NewDecoder(inputBytes)
	c.Assert(dec.Decode(s.pdu), IsNil)

	c.Check(int(s.pdu.HeaderLength), Equals, 34)
	c.Check(int(s.pdu.ApplicationId), Equals, mms.PUSH_APPLICATION_ID)
	c.Check(s.pdu.ContentType, Equals, mms.VND_WAP_MMS_MESSAGE)
	c.Check(len(s.pdu.Data), Equals, 102)
}
