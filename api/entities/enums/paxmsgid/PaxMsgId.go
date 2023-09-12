package paxmsgid

type PaxMsgId string

const (
	A00_INITIALIZE                      PaxMsgId = "A00"
	A02_GET_VARIABLE                    PaxMsgId = "A02"
	A04_SET_VARIABLE                    PaxMsgId = "A04"
	A06_SHOW_DIALOG                     PaxMsgId = "A06"
	A08_GET_SIGNATURE                   PaxMsgId = "A08"
	A10_SHOW_MESSAGE                    PaxMsgId = "A10"
	A12_CLEAR_MESSAGE                   PaxMsgId = "A12"
	A14_CANCEL                          PaxMsgId = "A14"
	A16_RESET                           PaxMsgId = "A16"
	A18_UPDATE_RESOURCE_FILE            PaxMsgId = "A18"
	A20_DO_SIGNATURE                    PaxMsgId = "A20"
	A22_DELETE_IMAGE                    PaxMsgId = "A22"
	A24_SHOW_MESSAGE_CENTER_ALIGNED     PaxMsgId = "A24"
	A26_REBOOT                          PaxMsgId = "A26"
	A28_GET_PIN_BLOCK                   PaxMsgId = "A28"
	A30_INPUT_ACCOUNT                   PaxMsgId = "A30"
	A32_RESET_MSR                       PaxMsgId = "A32"
	A36_INPUT_TEXT                      PaxMsgId = "A36"
	A38_CHECK_FILE                      PaxMsgId = "A38"
	A40_AUTHORIZE_CARD                  PaxMsgId = "A40"
	A42_COMPLETE_ONLINE_EMV             PaxMsgId = "A42"
	A44_REMOVE_CARD                     PaxMsgId = "A44"
	A46_GET_EMV_TLV_DATA                PaxMsgId = "A46"
	A48_SET_EMV_TLV_DATA                PaxMsgId = "A48"
	A50_INPUT_ACCOUNT_WITH_EMV          PaxMsgId = "A50"
	A52_COMPLETE_CONTACTLESS_EMV        PaxMsgId = "A52"
	A54_SET_SAF_PARAMETERS              PaxMsgId = "A54"
	A56_SHOW_TEXTBOX                    PaxMsgId = "A56"
	T00_DO_CREDIT                       PaxMsgId = "T00"
	T02_DO_DEBIT                        PaxMsgId = "T02"
	T04_DO_EBT                          PaxMsgId = "T04"
	T06_DO_GIFT                         PaxMsgId = "T06"
	T08_DO_LOYALTY                      PaxMsgId = "T08"
	T10_DO_CASH                         PaxMsgId = "T10"
	T12_DO_CHECK                        PaxMsgId = "T12"
	B00_BATCH_CLOSE                     PaxMsgId = "B00"
	B02_FORCE_BATCH_CLOSE               PaxMsgId = "B02"
	B04_BATCH_CLEAR                     PaxMsgId = "B04"
	B06_PURGE_BATCH                     PaxMsgId = "B06"
	B08_SAF_UPLOAD                      PaxMsgId = "B08"
	B10_DELETE_SAF_FILE                 PaxMsgId = "B10"
	R00_LOCAL_TOTAL_REPORT              PaxMsgId = "R00"
	R02_LOCAL_DETAIL_REPORT             PaxMsgId = "R02"
	R04_LOCAL_FAILED_REPORT             PaxMsgId = "R04"
	R06_HOST_REPORT                     PaxMsgId = "R06"
	R08_HISTORY_REPORT                  PaxMsgId = "R08"
	R10_SAF_SUMMARY_REPORT              PaxMsgId = "R10"
	A01_RSP_INITIALIZE                  PaxMsgId = "A01"
	A03_RSP_GET_VARIABLE                PaxMsgId = "A03"
	A05_RSP_SET_VARIABLE                PaxMsgId = "A05"
	A07_RSP_SHOW_DIALOG                 PaxMsgId = "A07"
	A09_RSP_GET_SIGNATURE               PaxMsgId = "A09"
	A11_RSP_SHOW_MESSAGE                PaxMsgId = "A11"
	A13_RSP_CLEAR_MESSAGE               PaxMsgId = "A13"
	A17_RSP_RESET                       PaxMsgId = "A17"
	A19_RSP_UPDATE_RESOURCE_FILE        PaxMsgId = "A19"
	A21_RSP_DO_SIGNATURE                PaxMsgId = "A21"
	A23_RSP_DELETE_IMAGE                PaxMsgId = "A23"
	A25_RSP_SHOW_MESSAGE_CENTER_ALIGNED PaxMsgId = "A25"
	A27_RSP_REBOOT                      PaxMsgId = "A27"
	A29_RSP_GET_PIN_BLOCK               PaxMsgId = "A29"
	A31_RSP_INPUT_ACCOUNT               PaxMsgId = "A31"
	A33_RSP_RESET_MSR                   PaxMsgId = "A33"
	A35_RSP_REPORT_STATUS               PaxMsgId = "A35"
	A37_RSP_INPUT_TEXT                  PaxMsgId = "A37"
	A38_RSP_CHECK_FILE                  PaxMsgId = "A39"
	A41_RSP_AUTHORIZE_CARD              PaxMsgId = "A41"
	A43_RSP_COMPLETE_ONLINE_EMV         PaxMsgId = "A43"
	A45_RSP_REMOVE_CARD                 PaxMsgId = "A45"
	A47_RSP_GET_EMV_TLV_DATA            PaxMsgId = "A47"
	A49_RSP_SET_EMV_TLV_DATA            PaxMsgId = "A49"
	A51_RSP_INPUT_ACCOUNT_WITH_EMV      PaxMsgId = "A51"
	A53_RSP_COMPLETE_CONTACTLESS_EMV    PaxMsgId = "A53"
	A55_RSP_SET_SAF_PARAMETERS          PaxMsgId = "A55"
	A57_RSP_SHOW_TEXTBOX                PaxMsgId = "A57"
	T01_RSP_DO_CREDIT                   PaxMsgId = "T01"
	T03_RSP_DO_DEBIT                    PaxMsgId = "T03"
	T05_RSP_DO_EBT                      PaxMsgId = "T05"
	T07_RSP_DO_GIFT                     PaxMsgId = "T07"
	T09_RSP_DO_LOYALTY                  PaxMsgId = "T09"
	T11_RSP_DO_CASH                     PaxMsgId = "T11"
	T13_RSP_DO_CHECK                    PaxMsgId = "T13"
	B01_RSP_BATCH_CLOSE                 PaxMsgId = "B01"
	B03_RSP_FORCE_BATCH_CLOSE           PaxMsgId = "B03"
	B05_RSP_BATCH_CLEAR                 PaxMsgId = "B05"
	B07_RSP_PURGE_BATCH                 PaxMsgId = "B07"
	B09_RSP_SAF_UPLOAD                  PaxMsgId = "B09"
	B11_RSP_DELETE_SAF_FILE             PaxMsgId = "B11"
	R01_RSP_LOCAL_TOTAL_REPORT          PaxMsgId = "R01"
	R03_RSP_LOCAL_DETAIL_REPORT         PaxMsgId = "R03"
	R05_RSP_LOCAL_FAILED_REPORT         PaxMsgId = "R05"
	R07_RSP_HOST_REPORT                 PaxMsgId = "R07"
	R09_RSP_HISTORY_REPORT              PaxMsgId = "R09"
	R11_RSP_SAF_SUMMARY_REPORT          PaxMsgId = "R11"
)

func (p PaxMsgId) GetValue() string {
	return string(p)
}

func (p PaxMsgId) GetBytes() []byte {
	return []byte(p)
}