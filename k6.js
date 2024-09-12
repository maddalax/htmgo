import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 100 }, // Ramp-up to 100 RPS over 1 minute
        { duration: '10m', target: 100 }, // Stay at 100 RPS for 10 minutes
        { duration: '1m', target: 0 },    // Ramp-down to 0 RPS
    ],
    thresholds: {
        http_req_duration: ['p(95)<500'], // 95% of requests should be below 500ms
    },
};

export default function () {
    http.get('http://localhost:3000/patients');
    sleep(1 / 100); // Make 100 requests per second
}
