const fileInput = document.getElementById("file-to-corrupt");

fileInput.addEventListener("change", async () => {
    const files = fileInput.files;

    if (!files || files.length === 0) {
        return;
    }

    const formData = new FormData();

    for (const file of files) {
        formData.append("file", file);
    }

    try {
        const response = await fetch("http://localhost:8080/file-corruptor/api", {
            method: "POST",
            body: formData
        });

        if (!response.ok) {
            throw new Error(`Request failed: ${response.status}`);
        }

        console.log("Files uploaded successfully");
    } catch (error) {
        console.error("Upload failed:", error);
    }
});