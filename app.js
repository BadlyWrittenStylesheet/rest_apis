const number1 = document.getElementById('num1')
const number2 = document.getElementById('num2')
const sendButton = document.getElementById('send')
const resultElem = document.getElementById('result')

async function callApi(val1, val2) {
    const response = await fetch('http://localhost:8080/api', {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            num1: val1,
            num2: val2
        })
    });
    return response.json()
}


async function manageClick() {
    if (!(number1.value && number2.value)) {
        resultElem.innerText = 'Input in both';
    }
    const response = await callApi(Number(number1.value), Number(number2.value));
        // .then(response => {
    console.log('repose', response);
    resultElem.innerText = response.result;
            // }
        // )
    // })
}

sendButton.addEventListener('click', manageClick)