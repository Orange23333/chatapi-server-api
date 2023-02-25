package stamping

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type StampingHandler struct {
	subHandler http.Handler
}

func New(subHandler http.Handler) *StampingHandler {
	return &StampingHandler{
		subHandler: subHandler,
	}
}

func (h *StampingHandler) SetSubHandler(subHandler http.Handler) {
	h.subHandler = subHandler
}

const HEADER_KEY_TIME_STAMP string = "date"
const HEADER_KEY_RECEIVE_TIME_STAMP string = "received_date"
const URL_KEY_TIME_STAMP string = "time_stamp"
const TIME_STAMP_EXAMPLE string = "Mon, 2 Jan 2006 15:04:05 MST"
const DEFAULT_TIMEOUT_SEC = 10

func (h *StampingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Client sends `time_stamp` support package security.
	// If request lost `time_stamp`, the security is lower.

	if h.subHandler != nil {
		time_now := time.Now()
		loc, _ := time.LoadLocation("GMT")
		time_now = time_now.In(loc)

		r.Header.Set(HEADER_KEY_RECEIVE_TIME_STAMP, time_now.Format(TIME_STAMP_EXAMPLE))

		h.subHandler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(500)
	}
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func CheckTimeStamp(time_stamp time.Time, reference_time time.Time, time_out_sec int) bool {
	return abs(time_stamp.Unix()-reference_time.Unix()) <= int64(time_out_sec) //Unix() is int64!
}

func CheckHttpTimeStamp(time_out_sec int, w http.ResponseWriter, r *http.Request, ps httprouter.Params) (bool, error) {
	// !This AD/DA step will waste time. How direct pass receive_time?
	receive_time_str := w.Header().Get(HEADER_KEY_RECEIVE_TIME_STAMP)
	if receive_time_str == "" {
		return false, nil
	}

	var timeStampString string
	loc, _ := time.LoadLocation("GMT")

	receive_time, err := time.ParseInLocation(TIME_STAMP_EXAMPLE, receive_time_str, loc)
	if err != nil {
		return false, err
	}

	if ps != nil {
		timeStampString = ps.ByName(URL_KEY_TIME_STAMP)

		if timeStampString != "" {
			return checkTimeStampString(timeStampString, receive_time, time_out_sec, loc)
		}
	}

	timeStampString = r.Header.Get(HEADER_KEY_TIME_STAMP)
	if timeStampString != "" {
		return checkTimeStampString(timeStampString, receive_time, time_out_sec, loc)
	}

	return false, nil
}

func checkTimeStampString(timeStampString string, reference_time time.Time, time_out_sec int, loc *time.Location) (bool, error) {
	//var duration_zero time.Duration = 0
	t, err := time.ParseInLocation(TIME_STAMP_EXAMPLE, timeStampString, loc)

	if err != nil {
		return false, err
	}

	return CheckTimeStamp(t, reference_time, time_out_sec), nil
}

func TimoutHandler(time_out_sec int, w http.ResponseWriter, r *http.Request, ps httprouter.Params) bool {
	result, err := CheckHttpTimeStamp(time_out_sec, w, r, ps)

	if !result {
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500) //Maybe 4xx
		}
	}

	return result
}
