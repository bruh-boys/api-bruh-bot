// test with deno

const rawResponse = await fetch(
    "http://127.0.0.1:5000/email",
    {
        method: "POST",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            email: `aaa`,
        }),
    }
);
const content = await rawResponse.json();
console.log(content);

