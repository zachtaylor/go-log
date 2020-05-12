package log

import "ztaylor.me/cast"

const (
	nocolor = "\x1b[0m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	purple  = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
)

// Format controls the text output specifics
type Format struct {
	// TimeFmt is the time Format
	TimeFmt func(cast.Time) string
	// SourceFmt is the source length
	SourceFmt func(string) string
	// MessageFmt is the message length
	MessageFmt func(string) string
	// Colors is the colors to use per level, or nil when in no-color mode
	Colors map[Level]string
	// PathCut
	PathCut []string
}

// NewFormat creates Format
func NewFormat(tfmt func(cast.Time) string, sfmt, mfmt func(string) string, colors map[Level]string) *Format {
	return &Format{
		TimeFmt:    tfmt,
		SourceFmt:  sfmt,
		MessageFmt: mfmt,
		Colors:     colors,
		PathCut:    make([]string, 0),
	}
}

// CutPathSource adds the caller's package to the Format's PathCut
func (f *Format) CutPathSource() {
	f.CutPathWith(NewSource(1), 0)
}

// CutPathSourceParent adds the caller's n-recursive parent package to the Format's PathCut
func (f *Format) CutPathSourceParent(n int) {
	f.CutPathWith(NewSource(1), n)
}

// CutPathWith adds the calling source's path ancestor number to the path cutset
func (f *Format) CutPathWith(source *Source, parentno int) {
	filePath, _ := cast.SplitPath(source.File())
	for i := 0; cast.Contains(filePath, "/") && i <= parentno; i++ {
		filePath = filePath[:cast.LastIndex(filePath, "/")]
	}
	f.PathCut = append(f.PathCut, filePath)
	cast.SortSlice(f.PathCut, func(i, j int) bool {
		return f.PathCut[i] > f.PathCut[j]
	})
}

// FormatCutPath removes path components from a string
func (f *Format) FormatCutPath(str string) string {
	for _, path := range f.PathCut {
		if pkglen := len(path); pkglen > len(str)+1 {
		} else if prefix := str[:pkglen]; prefix != path {
		} else if len(str) > pkglen {
			if str[pkglen] == '/' {
				return str[pkglen+1:]
			}
		}
	}
	return str
}

// Format creates writable output
func (f *Format) Format(time cast.Time, src *Source, lvl Level, flds cast.JSON, msg string) []byte {
	var sb cast.StringBuilder
	sb.WriteString(f.TimeFmt(time))
	sb.WriteByte(32) // space
	if f.Colors != nil {
		sb.WriteString(f.Colors[lvl])
	} else {
		sb.WriteByte(lvl.ByteCode())
		sb.WriteByte(32) // space
	}
	sb.WriteString(f.SourceFmt(f.FormatCutPath(src.String())))
	sb.WriteByte(32) // space
	sb.WriteString(f.MessageFmt(msg))
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for _, k := range flds.GetKeys() {
		cast.Fprintf(&sb, " %s=%v", k, flds[k])
	}
	sb.WriteByte(10) // newline
	return []byte(sb.String())
}
