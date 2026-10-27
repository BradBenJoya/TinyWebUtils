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
        const response = await fetch("http://localhost:8000/file-corruptor/api", {
            method: "POST",
            body: formData
        });

        if (!response.ok) {
            throw new Error(`Request failed: ${response.status}`);
        }

        console.log("Files uploaded successfully");

        const blob = await response.blob();
        const url = URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = files[0].name + ".corrupted";
        a.click();
        URL.revokeObjectURL(url);
        console.log("Downloading file")
    } catch (error) {
        console.error("Upload failed:", error);
    }
});