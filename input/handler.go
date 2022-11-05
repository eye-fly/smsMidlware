package input

import (
	"net/http"
	"sms/out"
)

type backEnd struct {
	out *out.BackEnd
}

func NewBackEnd(out *out.BackEnd) *backEnd {
	return &backEnd{
		out: out,
	}
}

func (bc *backEnd) HandleSendSmsRequest(w http.ResponseWriter, r *http.Request) {

	stringIp := r.Header.Get(ipHeader)
	if len(stringIp) == 0 {
		writeResp(w, "idn't recive hostIP", http.StatusBadRequest)
		return
	}

	reciverNumber := r.Header.Get(reciverNumberHeader)
	if len(reciverNumber) == 0 {
		writeResp(w, "idn't recive reciver tel. number", http.StatusBadRequest)
		return
	}
	messageText := r.Header.Get(messageTextHeader)
	if len(messageText) == 0 {
		writeResp(w, "idn't recive message Text", http.StatusBadRequest)
		return
	}

	bc.out.SetHostIp(stringIp)
	err := bc.out.SendSms(reciverNumber, messageText)
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResp(w, "msg sent!", http.StatusOK)
}

func writeResp(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	w.Write([]byte("sending failed: " + err))
}
