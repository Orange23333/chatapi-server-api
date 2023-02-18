package stamping

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type StampingHandler struct {
	subHandler http.Handler

	// If it is set true, the timestamp will be as same as receive time stamp if the time stamp is empty or wrong.
	allowNoTimeStampRequest bool
}

func New(subHandler http.Handler, allowNoTimeStampRequest bool) *StampingHandler {
	return &StampingHandler{
		subHandler:              subHandler,
		allowNoTimeStampRequest: allowNoTimeStampRequest,
	}
}

func (h *StampingHandler) SetSubHandler(subHandler http.Handler) {
	h.subHandler = subHandler
}

const HEADER_KEY_TIME_STAMP string = "date"
const HEADER_KEY_RECEIVE_TIME_STAMP string = "received_date"
const URL_KEY_TIME_STAMP string = "time_stamp"
const TIME_STAMP_EXAMPLE string = "Mon, 2 Jan 2006 15:04:05 MST"

func (h *StampingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Client sends `time_stamp` support package security. If request lost `time_stamp`, the security is lower.

	loc, _ := time.LoadLocation("GMT")
	time_now := time.Now().In(loc)
	var duration_zero time.Duration = 0

	time_stamp_str := r.Header.Get(HEADER_KEY_TIME_STAMP)
	t, _ := time.ParseInLocation(TIME_STAMP_EXAMPLE, time_stamp_str, loc)

	if (time_stamp_str == "") || (t.Sub(time_now) > duration_zero) {
		r.Header.Set(HEADER_KEY_TIME_STAMP, time_now.Format(TIME_STAMP_EXAMPLE))
	}

	w.Header().Add(HEADER_KEY_RECEIVE_TIME_STAMP, time_now.Format(TIME_STAMP_EXAMPLE))

	if h.subHandler != nil {
		h.subHandler.ServeHTTP(w, r)
	}
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func CheckTimeStamp(time_stamp time.Time, reference_time time.Time, time_out_sec int) bool {
	return Abs(time_stamp.Unix()-reference_time.Unix()) <= int64(time_out_sec) //Unix() is int64!
}

func CheckHttpTimeStamp(reference_time time.Time, time_out_sec int, w http.ResponseWriter, r *http.Request, ps httprouter.Params) (bool, error) {
	var timeStampString string

	loc, _ := time.LoadLocation("GMT")

	if ps != nil {
		timeStampString = ps.ByName(URL_KEY_TIME_STAMP)

		if timeStampString != "" {
			return checkTimeStampString(timeStampString, reference_time, time_out_sec, loc)
		}
	}

	timeStampString = r.Header.Get(HEADER_KEY_TIME_STAMP)
	if timeStampString != "" {
		return checkTimeStampString(timeStampString, reference_time, time_out_sec, loc)
	}

	return false, nil
}

func checkTimeStampString(timeStampString string, reference_time time.Time, time_out_sec int, loc *time.Location) (bool, error) {
	t, err := time.ParseInLocation(TIME_STAMP_EXAMPLE, timeStampString, loc)

	if err != nil {
		return false, err
	}

	return CheckTimeStamp(t, reference_time, time_out_sec), nil
}
