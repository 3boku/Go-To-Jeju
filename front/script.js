const postMessage = async () => {
    try {
        const input = document.querySelector(".input");
        const message = input.value;
        const feild = document.querySelector(".text3");

        if (message.trim() === "") {
            alert("Please enter a message");
            return;
        }

        feild.innerHTML = "답변을 준비중입니다...";

        input.value = "";

        // Send the message to the server
        const result = await fetch("http://localhost:8080/chat", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },

            body: JSON.stringify({ message }),
        }).then((r) => r.json());

        console.log(result);

        feild.innerHTML = result.message;
    } catch (error) {
        console.error(error);
        feild.innerHTML = "키쿠리쨩 술먹어서 어무것도 기억안나";
    }
};
