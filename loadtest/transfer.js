import http from "k6/http";

export default function () {
    const from = () => `${Math.floor(Math.random() * 1e6) + 1}`.padStart(10, '0');
    const to = () => `${Math.floor(Math.random() * 1e6) + 1}`.padStart(10, '0');

    const payload = JSON.stringify({
        from_account_no: from(),
        to_account_no: to(),
        amount: 10000,
    });

    http.post("http://localhost:8080/api/transfer", payload);
}
