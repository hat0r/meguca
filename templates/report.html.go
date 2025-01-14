// Code generated by qtc from "report.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line report.html:1
package templates

//line report.html:1
import "strconv"

//line report.html:2
import "time"

//line report.html:3
import "github.com/bakape/meguca/lang"

//line report.html:4
import "github.com/bakape/meguca/auth"

//line report.html:5
import "github.com/bakape/meguca/common"

// Report submission form

//line report.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line report.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line report.html:8
func StreamReportForm(qw422016 *qt422016.Writer, id uint64) {
//line report.html:9
	ln := lang.Get()

//line report.html:9
	qw422016.N().S(`<input type=text name=target value="`)
//line report.html:10
	qw422016.N().S(strconv.FormatUint(id, 10))
//line report.html:10
	qw422016.N().S(`" hidden><input type=text name=reason placeholder="`)
//line report.html:11
	qw422016.N().S(ln.Common.UI["reason"])
//line report.html:11
	qw422016.N().S(`" maxlength="`)
//line report.html:11
	qw422016.N().D(common.MaxLenReason)
//line report.html:11
	qw422016.N().S(`"><br><label><input type=checkbox name=illegal>`)
//line report.html:15
	qw422016.N().S(ln.UI["illegal"])
//line report.html:15
	qw422016.N().S(`<br></label>`)
//line report.html:18
	streamcaptcha(qw422016, "all")
//line report.html:19
	streamsubmit(qw422016, true)
//line report.html:20
}

//line report.html:20
func WriteReportForm(qq422016 qtio422016.Writer, id uint64) {
//line report.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line report.html:20
	StreamReportForm(qw422016, id)
//line report.html:20
	qt422016.ReleaseWriter(qw422016)
//line report.html:20
}

//line report.html:20
func ReportForm(id uint64) string {
//line report.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line report.html:20
	WriteReportForm(qb422016, id)
//line report.html:20
	qs422016 := string(qb422016.B)
//line report.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line report.html:20
	return qs422016
//line report.html:20
}

// Render list of all reports on board

//line report.html:23
func StreamReportList(qw422016 *qt422016.Writer, reports []auth.Report) {
//line report.html:24
	streamtableStyle(qw422016)
//line report.html:24
	qw422016.N().S(`<table>`)
//line report.html:26
	streamtableHeaders(qw422016, "id", "post", "reason", "time")
//line report.html:27
	for _, r := range reports {
//line report.html:27
		qw422016.N().S(`<tr><td>`)
//line report.html:29
		qw422016.N().S(strconv.FormatUint(r.ID, 10))
//line report.html:29
		qw422016.N().S(`</td><td>`)
//line report.html:30
		streamstaticPostLink(qw422016, r.Target)
//line report.html:30
		qw422016.N().S(`</td><td>`)
//line report.html:31
		qw422016.E().S(r.Reason)
//line report.html:31
		qw422016.N().S(`</td><td>`)
//line report.html:32
		qw422016.E().S(r.Created.Format(time.UnixDate))
//line report.html:32
		qw422016.N().S(`</td></tr>`)
//line report.html:34
	}
//line report.html:34
	qw422016.N().S(`</table>`)
//line report.html:36
}

//line report.html:36
func WriteReportList(qq422016 qtio422016.Writer, reports []auth.Report) {
//line report.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line report.html:36
	StreamReportList(qw422016, reports)
//line report.html:36
	qt422016.ReleaseWriter(qw422016)
//line report.html:36
}

//line report.html:36
func ReportList(reports []auth.Report) string {
//line report.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line report.html:36
	WriteReportList(qb422016, reports)
//line report.html:36
	qs422016 := string(qb422016.B)
//line report.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line report.html:36
	return qs422016
//line report.html:36
}
