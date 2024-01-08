// Code generated by command: go run src.go -out ../amd64.s -stubs ../stubs_amd64.go -pkg common. DO NOT EDIT.

//go:build amd64

package common

//go:noescape
func nttAVX2(p *[256]uint32)

//go:noescape
func invNttAVX2(p *[256]uint32)

//go:noescape
func mulHatAVX2(p *[256]uint32, a *[256]uint32, b *[256]uint32)

//go:noescape
func addAVX2(p *[256]uint32, a *[256]uint32, b *[256]uint32)

//go:noescape
func subAVX2(p *[256]uint32, a *[256]uint32, b *[256]uint32)

//go:noescape
func packLe16AVX2(p *[256]uint32, buf *byte)

//go:noescape
func reduceLe2QAVX2(p *[256]uint32)

//go:noescape
func le2qModQAVX2(p *[256]uint32)

//go:noescape
func exceedsAVX2(p *[256]uint32, bound uint32) uint8

//go:noescape
func mulBy2toDAVX2(p *[256]uint32, q *[256]uint32)
