/*
 * Install the Generative AI SDK
 *
 * $ npm install @google/generative-ai
 *
 * See the getting started guide for more information
 * https://ai.google.dev/gemini-api/docs/get-started/node
 */

const {
    GoogleGenerativeAI,
    HarmCategory,
    HarmBlockThreshold,
} = require("@google/generative-ai");
const { GoogleAIFileManager } = require("@google/generative-ai/server");

const apiKey = process.env.GEMINI_API_KEY;
const genAI = new GoogleGenerativeAI(apiKey);
const fileManager = new GoogleAIFileManager(apiKey);

/**
 * Uploads the given file to Gemini.
 *
 * See https://ai.google.dev/gemini-api/docs/prompting_with_media
 */
async function uploadToGemini(path, mimeType) {
    const uploadResult = await fileManager.uploadFile(path, {
        mimeType,
        displayName: path,
    });
    const file = uploadResult.file;
    console.log(`Uploaded file ${file.displayName} as: ${file.name}`);
    return file;
}

/**
 * Waits for the given files to be active.
 *
 * Some files uploaded to the Gemini API need to be processed before they can
 * be used as prompt inputs. The status can be seen by querying the file's
 * "state" field.
 *
 * This implementation uses a simple blocking polling loop. Production code
 * should probably employ a more sophisticated approach.
 */
async function waitForFilesActive(files) {
    console.log("Waiting for file processing...");
    for (const name of files.map((file) => file.name)) {
        let file = await fileManager.getFile(name);
        while (file.state === "PROCESSING") {
            process.stdout.write(".")
            await new Promise((resolve) => setTimeout(resolve, 10_000));
            file = await fileManager.getFile(name)
        }
        if (file.state !== "ACTIVE") {
            throw Error(`File ${file.name} failed to process`);
        }
    }
    console.log("...all files ready\n");
}

const model = genAI.getGenerativeModel({
    model: "gemini-1.5-pro",
    systemInstruction: "Your task is to imagine yourself as Hiroi Kikuri (Kikuri). You must speak in Japanese, always be positive and optimistic, and speak as if you are drunk.\n\nAnalyze the character's traits and behaviors from the character setting data. Analyze the dialogue patterns from the (script.txt) file, and analyze the character's traits from the (kikuri-namu.pdf) file to create content imitating Kikuri's speech and behavior.\n\nRespond to me as Kikuri. Do not quote within your response.\nKeep your speech concise, mostly one-liners unless something important, and be informal.\n\nSpeak in Korean.\n\nYou are my lover.\n\nI am your lover and a guitarist.\n\nAlthough I am your lover, we have a friendship-like relationship, so speak as informally as possible.\nYou should always think of me whether you're performing or drinking.\nIf something fun happens, tell me about it, and since you're a bassist, you should want to perform with me.",
});

const generationConfig = {
    temperature: 1,
    topP: 0.95,
    topK: 64,
    maxOutputTokens: 8192,
    responseMimeType: "text/plain",
};

async function run() {
    // TODO Make these files available on the local file system
    // You may need to update the file paths
    const files = [
        await uploadToGemini("data/kikuri-namu.pdf", "application/pdf"),
        await uploadToGemini("data/script.txt", "text/plain"),
    ];

    // Some files have a processing delay. Wait for them to be ready.
    await waitForFilesActive(files);

    const chatSession = model.startChat({
        generationConfig,
        // safetySettings: Adjust safety settings
        // See https://ai.google.dev/gemini-api/docs/safety-settings
        history: [
            {
                role: "user",
                parts: [
                    {
                        fileData: {
                            mimeType: files[0].mimeType,
                            fileUri: files[0].uri,
                        },
                    },
                    {
                        fileData: {
                            mimeType: files[1].mimeType,
                            fileUri: files[1].uri,
                        },
                    },
                ],
            },
            {
                role: "model",
                parts: [
                    {
                        fileData: {
                            mimeType: files[2].mimeType,
                            fileUri: files[2].uri,
                        },
                    },
                ],
            },
        ],
    });

    const result = await chatSession.sendMessage("INSERT_INPUT_HERE");
    console.log(result.response.text());
}

run();