// Code generated by qtc from "tag_values_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Tags generates response for /tags/<tag_name> handlerSee https://graphite.readthedocs.io/en/stable/tags.html#exploring-tags

//line app/vmselect/graphite/tag_values_response.qtpl:5
package graphite

//line app/vmselect/graphite/tag_values_response.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/graphite/tag_values_response.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/graphite/tag_values_response.qtpl:5
func StreamTagValuesResponse(qw422016 *qt422016.Writer, tagName string, tagValues []string) {
//line app/vmselect/graphite/tag_values_response.qtpl:5
	qw422016.N().S(`{"tag":`)
//line app/vmselect/graphite/tag_values_response.qtpl:7
	qw422016.N().Q(tagName)
//line app/vmselect/graphite/tag_values_response.qtpl:7
	qw422016.N().S(`,"values":[`)
//line app/vmselect/graphite/tag_values_response.qtpl:9
	for i, value := range tagValues {
//line app/vmselect/graphite/tag_values_response.qtpl:9
		qw422016.N().S(`{"count":1,"value":`)
//line app/vmselect/graphite/tag_values_response.qtpl:12
		qw422016.N().Q(value)
//line app/vmselect/graphite/tag_values_response.qtpl:12
		qw422016.N().S(`}`)
//line app/vmselect/graphite/tag_values_response.qtpl:14
		if i+1 < len(tagValues) {
//line app/vmselect/graphite/tag_values_response.qtpl:14
			qw422016.N().S(`,`)
//line app/vmselect/graphite/tag_values_response.qtpl:14
		}
//line app/vmselect/graphite/tag_values_response.qtpl:15
	}
//line app/vmselect/graphite/tag_values_response.qtpl:15
	qw422016.N().S(`]}`)
//line app/vmselect/graphite/tag_values_response.qtpl:18
}

//line app/vmselect/graphite/tag_values_response.qtpl:18
func WriteTagValuesResponse(qq422016 qtio422016.Writer, tagName string, tagValues []string) {
//line app/vmselect/graphite/tag_values_response.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/tag_values_response.qtpl:18
	StreamTagValuesResponse(qw422016, tagName, tagValues)
//line app/vmselect/graphite/tag_values_response.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/tag_values_response.qtpl:18
}

//line app/vmselect/graphite/tag_values_response.qtpl:18
func TagValuesResponse(tagName string, tagValues []string) string {
//line app/vmselect/graphite/tag_values_response.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/tag_values_response.qtpl:18
	WriteTagValuesResponse(qb422016, tagName, tagValues)
//line app/vmselect/graphite/tag_values_response.qtpl:18
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/tag_values_response.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/tag_values_response.qtpl:18
	return qs422016
//line app/vmselect/graphite/tag_values_response.qtpl:18
}