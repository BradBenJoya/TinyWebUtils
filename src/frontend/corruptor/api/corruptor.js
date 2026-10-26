const fileInput = document.getElementById("file-to-corrupt");

fileInput.addEventListener("change", async () => {
    const file = fileInput.files[0];

    if (!file) {
        return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
        const response = await fetch("http://localhost:8080/file-corruptor", {
            method: "POST",
            body: formData
        });

        if (!response.ok) {
            throw new Error(`Request failed: ${response.status}`);
        }

        console.log("File uploaded successfully");
    } catch (error) {
        console.error("Upload failed:", error);
    }
});