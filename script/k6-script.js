import http from 'k6/http';
import { check } from 'k6';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';

export const options = {
  batch: 20,
};

const createAccountBody = (documentNumber) => ({
  'document_number': documentNumber
});

const createTransactionBody = (accountId) => ({
    'account_id': accountId,
    'operation_type_id': 1,
    'amount': -100000
})

const appJsonHeaders = { headers: { 'Content-Type': 'application/json' }};
const host = 'http://localhost:8080';

export function setup() {
    const documentNumber = randomIntBetween(100000, 200000);
    const reqBody = createAccountBody(documentNumber);
    const reqBodyString = JSON.stringify(reqBody);
    console.log(`Account Creation Request body: ${reqBodyString}`);

    const res = http.post(`${host}/accounts`, reqBodyString, appJsonHeaders);
    check(res, {
      'status is 200': (r) => r.status === 200,
    });
    console.log(`Account Creation Response body: ${res.body}`);
    
    return { accountId: JSON.parse(res.body).id }
  }
  
  export default function(data) {
    const reqBody = createTransactionBody(data.accountId);
    const reqBodyString = JSON.stringify(reqBody);
    console.log(`Transaction Creation Request body: ${reqBodyString}`);
    const res = http.post(`${host}/transactions`, reqBodyString, appJsonHeaders);

  check(res, {
    'status is 200': (r) => r.status === 200,
    'has correct account id': (r) => r.json().account_id === reqBody.account_id,
    'has correct operation type id': (r) => r.json().operation_type_id === reqBody.operation_type_id,
    'has correct amount': (r) => r.json().amount === reqBody.amount,
  });

  console.log(`Transaction Creation Response body: ${res.body}`);
}