import http from "k6/http";

export default function () {
    const generateAccountNo = () => `${Math.floor(Math.random() * 1e6) + 1}`.padStart(10, '0');

    const payload = JSON.stringify({
        to_account_no: generateAccountNo(),
        amount: 1000000,
    });

    http.post("http://localhost:8080/api/transfer-in", payload);
}