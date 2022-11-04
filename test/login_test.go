package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestLogin(t *testing.T) {
	godotenv.Load("../.env")

	url := "https://www.airbnb.co.kr/authenticate"
	method := "POST"

	s := "email=%v&password=%v&remember_me=true&from=email_login&origin_url=https://www.airbnb.co.kr/?has_logged_out=1&page_controller_action_pair="
	s = fmt.Sprintf(s, os.Getenv("AIRBNB_EMAIL"), os.Getenv("AIRBNB_PASSWORD"))
	payload := strings.NewReader(s)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "www.airbnb.co.kr")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "180")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("X-Airbnb-Supports-Airlock-V2", "true")
	req.Header.Add("X-CSRF-Token", "V4$.airbnb.co.kr$gsxkPU--e8Y$9Svxc9jz6pVHHig84LTrtBHIgknRMUR6uVREh7haW6s=")
	req.Header.Add("X-Airbnb-Client-Action-ID", "719107f4-e147-469f-891a-554e88c548ea")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("dpr", "1")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("device-memory", "8")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("X-CSRF-Without-Token", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Add("viewport-width", "1146")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("ect", "4g")
	req.Header.Add("Origin", "https://www.airbnb.co.kr")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://www.airbnb.co.kr/?has_logged_out=1")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "_aat=0%7C8dL9gd9%2FvR00Xx9b6cOARtk9EIlBV6nj4KbAkfwfXXTqMdiWFQZVnuWUsQRQryjx; _airbed_session_id=c600195488680c47195045d914ef2b20; _csrf_token=V4%24.airbnb.co.kr%24yEeYcj6Prjc%24PLDTRFCzYZESoSUzK_BwS8nxs6ZdiGyauF_RzpXOXgQ%3D; _user_attributes=%7B%22curr%22%3A%22KRW%22%2C%22guest_exchange%22%3A1418.2199649999998%2C%22device_profiling_session_id%22%3A%221667365480--45174bcc2d757c101cd5a216%22%2C%22giftcard_profiling_session_id%22%3A%221667495828-265230538-6b59a93d0a628f00e98e91a5%22%2C%22reservation_profiling_session_id%22%3A%221667495828-265230538-0b665767fd00e90eac92c507%22%2C%22id%22%3A265230538%2C%22hash_user_id%22%3A%22e58516b717cb79f33165e6bbbf378455e50dee85%22%2C%22eid%22%3A%22KuPMZ2qZmrE0H_O_cKQnIQ%3D%3D%22%2C%22num_h%22%3A1%2C%22num_trip_notif%22%3A0%2C%22name%22%3A%22Sean%22%2C%22num_action%22%3A0%2C%22is_admin%22%3Afalse%2C%22can_access_photography%22%3Afalse%2C%22travel_credit_status%22%3Anull%2C%22referrals_info%22%3A%7B%22receiver_max_savings%22%3Anull%2C%22receiver_savings_percent%22%3Anull%2C%22receiver_signup%22%3Anull%2C%22referrer_guest%22%3A%22%E2%82%A916%2C000%22%2C%22terms_and_conditions_link%22%3A%22%2Fhelp%2Farticle%2F2269%22%2C%22wechat_link%22%3Anull%2C%22offer_discount_type%22%3Anull%7D%7D; abb_fa2=%7B%22user_id%22%3A%2268%7C1%7CE8zL%2FyuJ66xJW%2BxUMEldzjLQCpfSHGopBy%2B9MwUJSut3B0q9xpXSC6w%3D%22%7D; bev=1667365447_YThkMzQ4ZDNmNjdm; everest_cookie=1667495826.3g1YDSQ3G6p0BHrIFPaD.1vKdOtwUE63u7o7GlefWE-2iEyb0BzLwklfj-GclpYE; flags=1187844; jitney_client_session_created_at=1667495826; jitney_client_session_id=1e8d4172-a9a7-4a50-9453-e3330317285a; jitney_client_session_updated_at=1667495826; roles=0")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	// fmt.Println(res.Header)
	// headerFile, err := os.Create("data/headerFile")
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// headerFileDat, err := json.Marshal(headerFile)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// headerFile.Write(headerFileDat)
	// defer headerFile.Close()

	// fmt.Println(res.Cookies())

}
