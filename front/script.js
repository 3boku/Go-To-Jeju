const inputField = document.querySelector(".input");
const sendButton = document.querySelector(".send-btn");
const chatWrap = document.querySelector(".chatwrap");

sendButton.addEventListener("click", async () => {
    if (inputField.value.trim() !== "") {
        const userMessage = inputField.value;
        inputField.value = "";

        // 사용자 메시지 화면에 추가
        appendMessage("user", userMessage);

        // 서버 응답 대기 메시지 표시
        const thinkingMessage = appendMessage("bot", "키쿠리가 고민중...");

        try {
            const response = await fetch("http://localhost:8080/chat", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ message: userMessage }),
            });

            const data = await response.json();
            thinkingMessage.textContent = data.message;
        } catch (error) {
            console.error("Error:", error);
            thinkingMessage.textContent = "죄송합니다. 오류가 발생했습니다.";
        }
    }
});

function appendMessage(sender, message) {
    const messageElement = document.createElement("div");
    messageElement.classList.add("text3", sender);
    messageElement.textContent = message;
    chatWrap.appendChild(messageElement);
    chatWrap.scrollTop = chatWrap.scrollHeight;
    return messageElement;
}
